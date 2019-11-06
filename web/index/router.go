package index

import "github.com/go-chi/chi"

func Router() chi.Router {
	router := chi.NewRouter()

	// 主页显示文章列表, 包括: 文章名, 作者, 文章摘要
	router.Get("/", Index)

	return router
}
