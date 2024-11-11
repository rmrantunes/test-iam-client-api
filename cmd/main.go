package main

import (
	"sispa-iam-api/internal/config"
	"sispa-iam-api/internal/server"
)

func main() {
	casdoorClient := config.NewCasdoorClient()
	s := server.NewServer(casdoorClient)
	s.Start()
}
