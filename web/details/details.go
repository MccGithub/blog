package details

import (
	"github.com/MccGithub/blog/internal/dao"
	"github.com/MccGithub/blog/util"
	"github.com/MccGithub/blog/web/index"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Details(w http.ResponseWriter, r *http.Request) {
	tmplFiles := []string{
		"html/index.html",
		"html/articles.html",
		"html/details.html",
	}
	//id := r.FormValue("id")
	id := chi.URLParam(r, "article_id")
	db := r.Context().Value("db").(*dao.SQLHelper)
	article := db.GetArticle(id)
	if err := article.Get(); err != nil {
		index.Articles(w, r)
		return
	}
	data := util.Data{
		"target": "details",
		"details": article,
	}
	if err := util.View(w, data, tmplFiles...); err != nil {
		logrus.Warn(err)
	}
}
