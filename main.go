package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"websocketsProject/app"
	"websocketsProject/chat"
	"websocketsProject/controllers"
)

func main() {
	log.SetFlags(log.Lshortfile)

	router := mux.NewRouter()

	// websocket server
	server := chat.NewServer("/ws")
	go server.Listen(router)

	// SPA handler
	spa := app.SpaHandler{StaticPath: "js/build", IndexPath: "index.html"}
	router.PathPrefix("/app").Handler(spa)
	// static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("js/build/static"))))

	// api methods
	router.HandleFunc("/api/user", controllers.Current).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user/store", controllers.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/chats", controllers.ChatsOfCurrentUser).Methods("GET", "OPTIONS")


	// middlewares
	router.Use(app.AddCorsHeaders)
	router.Use(app.JwtAuthentication)
	router.Use(mux.CORSMethodMiddleware(router))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}


	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
