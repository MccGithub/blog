package details

import "github.com/go-chi/chi"

func Router() chi.Router {
	router := chi.NewRouter()

	// 文章详情页, 显示文章名, 作者, 内容
	router.Get("/", Details)

	return router
}
