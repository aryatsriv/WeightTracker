package main

import (
	"backend/internal/router"
	"log"
)

func main() {
	r := router.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
