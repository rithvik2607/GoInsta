package main

import (
	"net/http"

	"github.com/rithvik2607/GoInsta/config"
	"github.com/rithvik2607/GoInsta/posts"
	"github.com/rithvik2607/GoInsta/users"
)

func main() {
	// Connecting Database
	config.Connect()

	// Routes
	http.HandleFunc("/users", users.CreateUser)
	http.HandleFunc("/users/", users.GetUser)
	http.HandleFunc("/posts", posts.CreatePost)
	http.HandleFunc("/posts/", posts.GetPost)
	http.HandleFunc("/posts/users/", posts.GetUsersPost)

	// Server starts listening
	http.ListenAndServe(":3000", nil)
}
