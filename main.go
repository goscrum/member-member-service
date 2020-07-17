package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/goscrum/member-member-service/members"
	"net/http"
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

	err := http.ListenAndServe(":3333", r)
	if err != nil {
		panic([]byte("Some shit happened"))
	}
}
