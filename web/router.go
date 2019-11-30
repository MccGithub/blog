package web

import (
	"fmt"
	"net/http"

	"github.com/chive-chan/blog/internal/dao"
	"github.com/chive-chan/blog/util"
	"github.com/chive-chan/blog/web/author"
	"github.com/chive-chan/blog/web/index"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type Opt struct {
	Address string

	DBDriver string
	DBConn   string
}

func Serve(opt Opt) error {
	router := chi.NewRouter()
	//os.OpenFile()

	logrus.Trace("set db")
	router.Use(dao.DBHandler(opt.DBDriver, util.GetDbPath(opt.DBConn)))

	logrus.Tracef("%+v", opt)
	router.Mount("/", index.Router())
	router.Mount("/{author}", author.Router())

	fmt.Println("Listening at ", opt.Address)
	return http.ListenAndServe(opt.Address, router)
}
