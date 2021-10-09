package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rithvik2607/GoInsta/posts"
	"github.com/rithvik2607/GoInsta/users"
)

// Testing the Get User function
func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/users//mxjgsapgatlmodzl", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(users.GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[
		{
			"id": "mxjgsapgatlmodzl",
			"name": "rithvik",
			"email": "rithvik@test.com",
			"password": "GnutSYLOfYHOv4CdtwGlPaOIWytMz_K56uxnueTOVoMtRDP56R2Dm9VBDlbgsVXv37-iE1uUZkuftem-Z6v5Kg=="
		}
	]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// Testing the Create User function
func TestCreateUser(t *testing.T) {
	var jsonStr = []byte(`{"name": "Aryan", "email": "aryan@test.com", "password": "test12345"}`)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(users.CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{
		"id": "kwmhjxqb",
		"name": "Aryan",
		"email": "aryan@test.com",
		"password": "GnutSYLOfYHOv4CdtwGlPaOIWytMz_K56uxnueTOVoMtRDP56R2Dm9VBDlbgsVXv37-iE1uUZkuftem-Z6v5Kg=="
	}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// Testing the Create Post function
func TestCreatePost(t *testing.T) {
	var jsonStr = []byte(`{"user_id": "mxjgsapgatlmodzl", "caption": "Let's see if this works", "img_url": "https://www.test.com/image2"}`)

	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(posts.CreatePost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{
		"id": "ba1agqvf",
		"user_id": "mxjgsapgatlmodzl",
		"caption": "Let's see if this works",
		"img_url": "https://www.test.com/image2",
		"timestamp": "2021-10-09T20:12:49+05:30"
	}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// Testing the Get Post function
func TestGetPost(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/w3beddgl", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(posts.GetPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[
		{
			"id": "w3beddgl",
			"user_id": "mxjgsapgatlmodzl",
			"caption": "this is a test",
			"img_url": "https://test.com/img",
			"timestamp": "2021-10-09T17:09:47+05:30"
		}
	]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// Testing the Get User's Posts function
func TestGetUsersPost(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts/users/mxjgsapgatlmodzl", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(posts.GetUsersPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[
		{
			"id": "askjckjeb87368726176",
			"user_id": "mxjgsapgatlmodzl",
			"caption": "test message",
			"img_url": "https://test.com/image",
			"timestamp": "1997-07-16T19:20+01:00"
		},
		{
			"id": "w3beddgl",
			"user_id": "mxjgsapgatlmodzl",
			"caption": "this is a test",
			"img_url": "https://test.com/img",
			"timestamp": "2021-10-09T17:09:47+05:30"
		},
		{
			"id": "404d8he8",
			"user_id": "mxjgsapgatlmodzl",
			"caption": "Let's see if this works",
			"img_url": "https://www.test.com/image2",
			"timestamp": "2021-10-09T20:18:22+05:30"
		}
	]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
