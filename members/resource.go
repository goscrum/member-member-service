package members

import (
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
	r.Route("/{memberID}", func(sr chi.Router) {
		sr.Get("/", rs.GetByID)
	})

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
	ctx := r.Context()
	err := render.DecodeJSON(r.Body, &newMember)
	if err != nil {
		getError(err, http.StatusInternalServerError, w, r)
		return
	}
	newMember.ID = primitive.NewObjectID()
	dao := New(ctx)
	err = dao.Add(&newMember, ctx)
	if err != nil {
		getError(err, http.StatusBadRequest, w, r)
		return
	}

	render.JSON(w, r, newMember)
}

func (rs Resource) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	memberID := chi.URLParam(r, "memberID")
	oid, err := primitive.ObjectIDFromHex(memberID)
	if err != nil {
		getError(err, http.StatusNotFound, w, r)
	}

	dao := New(ctx)
	member, err := dao.GetByID(oid, ctx)
	if err != nil {
		getError(err, http.StatusNotFound, w, r)
		return
	}

	render.JSON(w, r, member)
}
