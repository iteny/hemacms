package router

import (
	"hemacms/ctrl/api"
	"net/http"

	"github.com/go-chi/chi"
)

func ApiRoutes() http.Handler {
	api := api.ApiCtrlObject()
	r := chi.NewRouter()
	r.Post("/api", api.GetSlider) //首页
	return r
}
