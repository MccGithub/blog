package article

import (
	"net/http"
	"net/url"

	"github.com/MccGithub/blog/internal/dao"
	"github.com/MccGithub/blog/util"
	"github.com/MccGithub/blog/web/index"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func Article(w http.ResponseWriter, r *http.Request) {
	tmplFiles := []string{
		"html/index.html",
		"html/articles.html",
		"html/details.html",
	}
	//id := r.FormValue("id")
	author_name, err := url.QueryUnescape(chi.URLParam(r, "author"))
	if err != nil {
		logrus.Warn(err)
	}
	article_name, err := url.QueryUnescape(chi.URLParam(r, "article"))
	if err != nil {
		logrus.Warn(err)
	}
	db := r.Context().Value("db").(*dao.SQLHelper)
	article := db.GetArticle(author_name, article_name)
	if err := article.Get(); err != nil {
		logrus.Warn(err)
		index.Index(w, r)
		return
	}
	//article.Content = strings.ReplaceAll(article.Content)
	data := util.Data{
		"target":  "details",
		"details": article,
	}
	if err := util.View(w, data, tmplFiles...); err != nil {
		logrus.Warn(err)
	}
}
