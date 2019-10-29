package main

import (
	"net/http"
	"server/routes"
)

func main() {
	routes.SetRouting()
	http.ListenAndServe(":8080", nil)
}