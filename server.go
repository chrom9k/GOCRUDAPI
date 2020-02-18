package main

import (
	"net/http"

	"github.com/chrom9k/GOCRUDAPI/routes"
)

func main() {
	routes.SetRouting()
	http.ListenAndServe(":8080", nil)
}
