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

	http.HandleFunc("/users", users.CreateUser)
	http.HandleFunc("/users/:id", users.GetUser)
	http.HandleFunc("/posts", posts.CreatePost)
	http.HandleFunc("/posts/:id", posts.GetPost)
	http.HandleFunc("/posts/users/:user_id", posts.GetUsersPost)
}
