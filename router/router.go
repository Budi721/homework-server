package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"homework-server/handler"
	"net/http"
)

func NewRouter(userHandler handler.UserHandler) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome Budi API :v")
	})
	r.Get("/{id}/detail", userHandler.GetDetail)
	r.Get("/follower/{username}", userHandler.GetFollowers)
	return r
}
