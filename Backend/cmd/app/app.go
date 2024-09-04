package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ApiServer struct {
	address string
	store   Storage
}

func NewServer(address string, store Storage) *ApiServer {
	return &ApiServer{
		address: address,
		store:   store,
	}
}

func (s *ApiServer) Run() {

	mux := http.NewServeMux()

	mux.HandleFunc("/user/", makeHTTPHandlFunc(s.handelUsers))
	// TODO other routs ...

	log.Print("Server running on port 8080 ... ")

	http.ListenAndServe(s.address, mux)
}

func (s *ApiServer) handelUsers(w http.ResponseWriter, r *http.Request) error {

	path := r.URL.Path

	if path == "/user/" {

		if r.Method == "GET" {
			log.Printf("Catch a  %s path with method %s", path, r.Method)

			return s.handelGetUsers(w, r)

		} else if r.Method == "POST" {
			log.Printf("Catch a  %s path with method %s", path, r.Method)

			return s.handelCreateUser(w, r)

		} else {

			return fmt.Errorf("method %s not allowed by url %s", r.Method, path)
		}

	} else {
		id, err := getUserIdFromPath(path)

		if err != nil {
			return err
		}

		if r.Method == "GET" {
			log.Printf("Catch a  %s path with method %s", path, r.Method)

			return s.handelGetUserById(w, r, id)

		} else if r.Method == "DELETE" {
			log.Printf("Catch a  %s path with method %s", path, r.Method)

			return s.handelDeleteUser(w, r, id)

		} else if r.Method == "PATCH" {
			log.Printf("Catch a  %s path with method %s", path, r.Method)

			return s.handelUpdateUser(w, r, id)

		} else {

			return fmt.Errorf("method %s not allowed by url %s", r.Method, path)
		}

	}

}

func (s *ApiServer) handelCreateUser(w http.ResponseWriter, r *http.Request) error {
	user := TransferUser{}

	json.NewDecoder(r.Body).Decode(&user)

	log.Printf("Got a new user %+v", user)

	err := s.store.CreateUser(&user)

	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, 0)
}

func (s *ApiServer) handelUpdateUser(w http.ResponseWriter, r *http.Request, id int) error {
	user := TransferUser{}

	json.NewDecoder(r.Body).Decode(&user)

	log.Printf("Got an update user name %+v with id = %d", user, id)

	err := s.store.UpdateUser(id, &user)

	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, 0)
}

func (s *ApiServer) handelDeleteUser(w http.ResponseWriter, r *http.Request, id int) error {

	err := s.store.DeleteUser(id)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, 0)
}

func (s *ApiServer) handelGetUsers(w http.ResponseWriter, r *http.Request) error {

	users, err := s.store.GetUsers()

	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, users)
}

func (s *ApiServer) handelGetUserById(w http.ResponseWriter, r *http.Request, id int) error {
	user, err := s.store.GetUserById(id)

	if err != nil {

		return err
	}

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
