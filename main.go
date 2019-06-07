package main

import (
	"fmt"
	"net/http"
	"os"

	"go-phonebooks/controllers"
	"go-phonebooks/middlewares"
	_ "go-phonebooks/utils/env"

	"regexp"

	"github.com/gorilla/mux"
)

func setMiddleware(middlewareName string, handlerFunc http.HandlerFunc) http.Handler {
	if middlewareName == "jwt" {
		return middlewares.JwtAuthMiddleware(handlerFunc)
	} else if middlewareName == "hello" {
		return middlewares.HelloMiddleware(handlerFunc)
	}
	return nil
}

func appendMiddleware(middlewareName string, otherMiddlewares http.Handler) http.Handler {
	if middlewareName == "jwt" {
		return middlewares.JwtAuthMiddleware(otherMiddlewares)
	} else if middlewareName == "hello" {
		return middlewares.HelloMiddleware(otherMiddlewares)
	}
	return otherMiddlewares
}

func recursiveMiddleware(myMiddleware http.Handler, middlewareArr []string, start int, handler http.HandlerFunc) http.Handler {
	if start < len(middlewareArr) {
		if myMiddleware == nil {
			myMiddleware = setMiddleware(middlewareArr[start], handler)
		} else {
			myMiddleware = appendMiddleware(middlewareArr[start], myMiddleware)
		}
		return recursiveMiddleware(myMiddleware, middlewareArr, start+1, handler)
	}
	return myMiddleware
}

func registerRoute(r *mux.Router, ctrl *controllers.Controller, name string, handler func(w http.ResponseWriter, r *http.Request)) {
	sRoute, ok := ctrl.Routes[name]
	myRoute := ctrl.PrefixURL
	regex, _ := regexp.Compile("/$")
	myRoute = regex.ReplaceAllString(myRoute, "")
	// fmt.Println(myRoute)
	// fmt.Println(sRoute)
	if sRoute.URL != "" {
		myRoute = myRoute + sRoute.URL
	}
	fmt.Println(myRoute)
	if !ok {
		fmt.Println("Cannot find index['" + name + "']")
		return
	}
	mids, ok := ctrl.Middlewares[name]
	if ok {
		handlerFunc := http.HandlerFunc(handler)
		enabledMiddlewares := []string{}
		var midFuncs http.Handler
		for _, value := range mids {
			if value == "jwt" || value == "hello" {
				enabledMiddlewares = append(enabledMiddlewares, value)
			}
		}
		if len(enabledMiddlewares) > 0 {
			// fmt.Println(enabledMiddlewares)
			midFuncs = recursiveMiddleware(midFuncs, enabledMiddlewares, 0, handlerFunc)
		}
		if midFuncs != nil {
			r.Handle(myRoute, midFuncs).
				Methods(sRoute.Method).
				Name(sRoute.Name)
		} else {
			r.HandleFunc(myRoute, handler).
				Methods(sRoute.Method).
				Name(sRoute.Name)
		}
	} else {
		r.HandleFunc(myRoute, handler).
			Methods(sRoute.Method).
			Name(sRoute.Name)
	}
}

func main() {
	router := mux.NewRouter()
	apiRoute := router.PathPrefix("/api").Subrouter()
	registerRoute(apiRoute, controllers.HomeController, "Index", controllers.HomeController.Index)
	registerRoute(apiRoute, controllers.HomeController, "Index2", controllers.HomeController.Index2)
	registerRoute(apiRoute, controllers.AuthController, "Profile", controllers.HomeController.Profile)
	registerRoute(apiRoute, controllers.AuthController, "Register", controllers.HomeController.Register)
	registerRoute(apiRoute, controllers.AuthController, "Login", controllers.HomeController.Login)
	// router.HandleFunc("/api", controllers.HomeController.Index).
	// 	Methods("GET")
	// router.Use(middlewares.JwtAuthMiddleware)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Your port: " + port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:5000/api
	if err != nil {
		fmt.Print(err)
	}
}
