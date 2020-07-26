package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/goscrum/member-member-service/members"
	"net/http"
	"os"
	"time"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("hello")); err != nil {
			panic(err)
		}
	})

	r.Mount("/members", members.Resource{}.Routes())

	serverPort := os.Getenv("APP_PORT")
	var err error
	if len(serverPort) > 0 {
		err = http.ListenAndServe(fmt.Sprintf(":%s", serverPort), r)
	} else {
		err = http.ListenAndServe(":8080", r)
	}
	if err != nil {
		panic([]byte("Some shit happened"))
	}
}
