package main

import (
	"log"

	"github.com/triangletodd/gort/internal/config"
	"github.com/triangletodd/gort/internal/server"
)

func main() {
	config.Init()
	log.Println("parsed config")
	server.Init()
	log.Println("Server init'd")
}
