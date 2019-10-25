package api

import "github.com/go-chi/chi"

func Router() chi.Router {
	router := chi.NewRouter()

	//router.Get("/xxx", Xxx)

	return router
}