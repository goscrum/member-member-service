package members

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

type Resource struct{}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Post("/", rs.Create)

	return r
}

func (rs Resource) Create(w http.ResponseWriter, r *http.Request) {
	var newMember Member
	err := render.DecodeJSON(r.Body, &newMember)
	if err != nil {
		render.Status(r, 500)
		render.NoContent(w, r)
	}
	render.JSON(w, r, newMember)
}
