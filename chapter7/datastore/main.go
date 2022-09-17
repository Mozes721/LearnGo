package main

import (
	"fmt" 
	"http"
)



type DataStore interface {
	UserNameForID(userID string) (string, bool)
}
type Logger interface {
	Log(message string)
}
type Logic interface {
	SayHello(userID string) (string, error)
}

type LoggerAdapeter func(message string)


func (lg LoggerAdapeter) Log(message string) {
	lg(message)
}

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

type SimpleLogic struct {
	l Logger
	ds DataStore

}

type Controler struct {
	l Logger
	logic Logic

}



func (c Controler) SayHello(w http.ResponseWritter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err := nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, logic Logic) Controler {
	return Controler{
		l: l,
		logic: logic,
	}
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if != ok {
		return "", error.New("unkown user")
	}
	return "Hello, " + name, nil	

}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in SayGoodbye for" + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if != ok {
		return "", error.New("unkown user")
	}
	return "Goodbye, " + name, nil	
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic {
		l: l,
		dsL ds
	}
}

func NewSimpleDataStore() SimpleDataStore {
	return simpleDataStore {
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat"
		},
	}
}

func main() {
	l := LoggerAdapeter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}