package members

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type Resource struct{}

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Post("/", rs.Create)

	return r
}

func getError(err error, status int, w http.ResponseWriter, r *http.Request) {
	log.Println(err.Error())
	response := ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   status,
	}

	render.JSON(w, r, response)
}

func (rs Resource) Create(w http.ResponseWriter, r *http.Request) {
	var newMember Member
	err := render.DecodeJSON(r.Body, &newMember)
	if err != nil {
		getError(err, http.StatusInternalServerError, w, r)
		return
	}
	newMember.ID = primitive.NewObjectID()
	dao := New()
	err = dao.Add(&newMember, context.TODO())
	if err != nil {
		getError(err, http.StatusBadRequest, w, r)
		return
	}

	render.JSON(w, r, newMember)
}
