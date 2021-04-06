package main

import "github.com/riberto-dalho/fullcycle-desafio/src/http"

func main() {
	server := http.NewWebserver()
	server.Serve(8000)
}
