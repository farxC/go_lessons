package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type api struct {
	addr string
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

var users = []User{}

// This is our server.
func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page"))
			return
		case "/users":
			w.Write([]byte("Users page!!!"))
			return
		default:
			w.Write([]byte("404 not found"))
			return
		}
	case http.MethodPost:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Not implemented yet\n"))
		}
	}
	w.Write([]byte("Hello from the server!"))
}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Age:       payload.Age,
	}
	if err = insertUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *api) helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func insertUser(u User) error {

	if u.FirstName == "" {
		return errors.New("first name is required")
	}
	if u.LastName == "" {
		return errors.New("last name is required")
	}
	if u.Age < 0 {
		return errors.New("age can't be minor than 0")
	}

	for _, user := range users {
		if user.Email == u.Email && user.FirstName == u.FirstName && user.LastName == u.LastName {
			return errors.New("User already exists")
		}
	}

	users = append(users, u)
	return nil
}

func main() {

	// Beautiful
	api := &api{
		addr: ":8080",
	}
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	//Router, beautiful

	// Routes: Both methods receive an reference from a any given variable that implements the "api" struct
	mux.HandleFunc("/hello", api.helloHandler)
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			api.getUsersHandler(w, r)
		case http.MethodPost:
			api.createUsersHandler(w, r)
		default:
			http.NotFound(w, r)
		}

	})

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
