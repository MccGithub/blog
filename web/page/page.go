package page

import (
	"github.com/MccGithub/blog/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if err := util.View(w, nil, "html/main.html", "html/article.html"); err != nil {
		logrus.Warn(err)
	}
}