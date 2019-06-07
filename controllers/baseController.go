package controllers

import (
	res "go-phonebooks/utils"
	"net/http"
)

type IController interface {
	GetPrefixUrl() string
	GetRoutes() map[string]Route
	GetMiddlewares() map[string][]string
}
type Controller struct {
	PrefixURL   string
	Routes      map[string]Route
	Middlewares map[string][]string
}
type HomeControllerType struct {
	Controller
}

type Route struct {
	URL    string
	Method string
	Name   string
}

func (i *HomeControllerType) GetPrefixUrl() string {
	return i.PrefixURL
}

func (i *HomeControllerType) GetRoutes() map[string]Route {
	return i.Routes
}
func (i *HomeControllerType) GetMiddlewares() map[string][]string {
	return i.Middlewares
}

var HomeController = &HomeControllerType{}

func init() {
	HomeController.PrefixURL = "/"
	routes := map[string]Route{
		"Index":  Route{URL: "", Method: http.MethodGet, Name: "Home.Get.Index"},
		"Index2": Route{URL: "/home", Method: http.MethodPost, Name: "Home.Post.Index2"},
	}
	mids := map[string][]string{"Index": []string{"jwt", "hello"}}
	HomeController.Middlewares = mids
	HomeController.Routes = routes
}
func (self *HomeControllerType) Index(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world!!!"
	res.Respond(w, 200, res.Message(200, msg))
}
func (self *HomeControllerType) Index2(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world!!! (w/o middlewares)"
	res.Respond(w, 200, res.Message(200, msg))
}
