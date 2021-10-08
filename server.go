package main

import (
	"net/http"

	"github.com/rithvik2607/GoInsta/users"
)

func main() {
	client = config.initDataLayer()

	http.HandleFunc("/users", users.GetAllUsers)
	http.HandleFunc("/users/:id")
}
