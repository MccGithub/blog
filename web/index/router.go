package page

import "github.com/go-chi/chi"

func Router() chi.Router {
	router := chi.NewRouter()

	router.Get("/", Article)
	router.Get("/article", Article)

	return router
}
