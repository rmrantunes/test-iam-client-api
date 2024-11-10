package main

import "sispa-iam-api/internal/server"

func main() {
	s := server.NewServer()
	s.Start()
}
