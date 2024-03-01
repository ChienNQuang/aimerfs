package main

import "fmt"

func main() {
	server := newServer()
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
