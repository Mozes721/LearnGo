package contexts

import (
	"context"
	"net/http"
)

type Logger interface {
	Log(context.Context, string)
}
type RequestDecorator func(*http.Request) *http.Request

type BusinessLogic struct {
	RequestDecorator RequestDecorator
	Logger           Logger
	remote           string
}

func (bl *BusinessLogic) BusinessLogic(ctx context.Context, user string, data string) (*http.Response, error) {
	bl.Logger.Log(ctx, "startign businessLogic for "+user+" with "+data)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, bl.remote+"?query="+data, nil)
	if err != nil {
		bl.Logger.Log(ctx, "error building remote request")
		return "", err
	}
	req = bl.RequestDecorator(req)
	resp, err := http.DefaultClient.Do(req)
	return resp, err
}
