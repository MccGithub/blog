package page

import "github.com/go-chi/chi"

func Router() chi.Router {
	router := chi.NewRouter()

	router.Get("/", Home)

	return router
}
