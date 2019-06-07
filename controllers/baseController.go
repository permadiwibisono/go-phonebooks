package controllers

import (
	res "go-phonebooks/utils"
	"net/http"
)

type Controller struct {
	PrefixURL   string
	Routes      map[string]Route
	Middlewares map[string][]string
}

type Route struct {
	URL    string
	Method string
	Name   string
}

var HomeController = &Controller{PrefixURL: "/"}

func init() {
	routes := map[string]Route{
		"Index":  Route{URL: "", Method: http.MethodGet, Name: "Home.Get.Index"},
		"Index2": Route{URL: "/home", Method: http.MethodPost, Name: "Home.Post.Index2"},
	}
	mids := map[string][]string{"Index": []string{"jwt", "hello"}}
	HomeController.Middlewares = mids
	HomeController.Routes = routes
}
func (self *Controller) Index(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world!!!"
	res.Respond(w, 200, res.Message(200, msg))
}
func (self *Controller) Index2(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world!!! (w/o middlewares)"
	res.Respond(w, 200, res.Message(200, msg))
}
