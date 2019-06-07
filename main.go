package main

import (
	"fmt"
	"net/http"
	"os"

	"go-phonebooks/bootstrap"
)

func main() {
	app := &bootstrap.App{}
	app.Initialize()
	// router := mux.NewRouter()
	// apiRoute := router.PathPrefix("/api").Subrouter()
	// settingControllerRoutes(apiRoute, controllers.HomeController)
	// settingControllerRoutes(apiRoute, controllers.AuthController)
	// registerRoute(apiRoute, controllers.HomeController, "Index", controllers.HomeController.Index)
	// registerRoute(apiRoute, controllers.HomeController, "Index2", controllers.HomeController.Index2)
	// registerRoute(apiRoute, controllers.AuthController, "Profile", controllers.AuthController.Profile)
	// registerRoute(apiRoute, controllers.AuthController, "Register", controllers.AuthController.Register)
	// registerRoute(apiRoute, controllers.AuthController, "Login", controllers.AuthController.Login)
	// router.HandleFunc("/api", controllers.HomeController.Index).
	// 	Methods("GET")
	// router.Use(middlewares.JwtAuthMiddleware)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Your port: " + port)

	err := http.ListenAndServe(":"+port, app.Router) //Launch the app, visit localhost:5000/api
	if err != nil {
		fmt.Print(err)
	}
}
