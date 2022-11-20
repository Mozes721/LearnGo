package cancelation

import (
	"context"
	"net/http"
)

type userKey int

const key userKey = 1

func ContextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, key, user)
}

func UserFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(key).(string)
	return user, ok
}
func extractUser(req *http.Request) (string, error) {
	userCookie, err := req.Cookie("user")
	if err != nil {
		return "", err
	}
	return userCookie.Value, nil
}
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		user, err := extractUser(req)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := req.Context()
		ctx = ContextWithUser(ctx, user)
		req = req.WithContext(ctx)
		h.ServeHTTP(rw, req)
	})
}
func (c Controller) handleRequest(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	user, ok := identity.UserFromContext(ctx)
	if !ok {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := req.URL.Query().Get("data")
	result, err := c.Logic.businessLogic(ctx, user, data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write([]byte(result))
}

func longRunningThingManger(ctx context.Context, data string) (string, error) {
	type wrapper struct {
		result string
		err    error
	}
	ch := make(chan wrapper, 1)
	go func() {
		result, err := longRunningThing(ctx, data)
		ch <- wrapper{result, err}
	}()
	select {
	case data = <-ch:
		return data.result, data.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
