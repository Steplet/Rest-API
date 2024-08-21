package app

import (
	"log"
	"net/http"
)

type ApiServer struct {
	Address string
}

func NewServer(address string) *ApiServer {
	return &ApiServer{Address: address}
}

func (s *ApiServer) Run() {
	log.Print("running... ")

	http.ListenAndServe(s.Address, nil)
}
