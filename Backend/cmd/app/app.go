package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func NewServer(address string) *ApiServer {
	return &ApiServer{Address: address}
}

func (s *ApiServer) Run() {

	mux := http.NewServeMux()

	mux.HandleFunc("/user/", makeHTTPHandlFunc(s.handelUsers))

	log.Print("Server running on port 8080 ... ")

	http.ListenAndServe(s.Address, mux)
}

func (s *ApiServer) handelUsers(w http.ResponseWriter, r *http.Request) error {

	path := r.URL.Path

	if r.Method == "GET" && path == "/user/" {
		log.Printf("Catch a  %s path", path)
		return s.handelGetUsers(w, r)
	}

	if r.Method == "GET" {

		id, err := getUserIdFromPath(path)

		if err != nil {
			return err
		}

		log.Printf("Catch a  %s path with id = %d", path, id)

		return s.handelGetUserById(w, r)
	}

	return fmt.Errorf("bad method %s", r.Method)

}

func (s *ApiServer) handelGetUsers(w http.ResponseWriter, r *http.Request) error {
	user1 := NewUser("Artem")
	user2 := NewUser("Ivan")
	users := []User{*user1, *user2}
	return WriteJSON(w, http.StatusOK, users)
}

func (s *ApiServer) handelGetUserById(w http.ResponseWriter, r *http.Request) error {
	user := NewUser("Man")
	return WriteJSON(w, http.StatusOK, user)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandlFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type ApiServer struct {
	Address string
}

func getUserIdFromPath(path string) (int, error) {

	pathParts := strings.Split(strings.Trim(path, "/"), "/")

	if len(pathParts) != 2 {
		return -1, fmt.Errorf("expected 2 args, but get %d", len(pathParts))
	}

	id, err := strconv.Atoi(pathParts[1])

	if err != nil {
		return -1, fmt.Errorf("can not convert string to id. bad args")
	}

	return id, nil
}
