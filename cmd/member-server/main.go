package main

import (
	"log"
	"net/http"

	"github.com/goscrum/member-member-service/pkg/http/rest"
)

func main() {
	handler := rest.SetupHandler()

	log.Fatal(http.ListenAndServe(":8080", handler))
}
