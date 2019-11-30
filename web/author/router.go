package author

import (
	"github.com/chive-chan/blog/web/author/article"
	"github.com/go-chi/chi"
)

func Router() chi.Router {
	router := chi.NewRouter()

	// 文章详情页, 显示文章名, 作者, 内容
	router.Mount("/{article}", article.Router())

	return router
}
