package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dybaay/go-server/routes"
)

func main() {
	routes.Routes()
	fmt.Printf("Starting server at port 8082\n")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}
