package main

import (
	"fmt"
	"net/http"
	"os"

	middlewares "go-phonebooks/middlewares"
	"go-phonebooks/models"
	res "go-phonebooks/utils"
	_ "go-phonebooks/utils/env"

	"github.com/gorilla/mux"
)

func HomeRouteHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world!!!"
	user := &models.User{
		Email: "mbapewe@gmail.com",
	}
	err, isValid := user.Validate()
	fmt.Println(err)
	fmt.Println(isValid)
	if err != nil {
		res.RespondError(w, 422, "Invalid request!", err)
		return
	}
	res.Respond(w, 200, res.Message(200, msg))
	// fmt.Println(msg)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeRouteHandler)
	router.Use(middlewares.JwtAuthMiddleware)
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
