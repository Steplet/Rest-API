package app

import "fmt"

type ApiServer struct {
	Address string
}

func (s ApiServer) NewServer(address string) *ApiServer {
	return &ApiServer{Address: address}
}

func TestModule() {
	fmt.Println("Hello from app")
}
