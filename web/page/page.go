package page

import (
	"github.com/MccGithub/blog/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if err := util.View("html/main.html", w, nil); err != nil {
		logrus.Warn(err)
	}
}