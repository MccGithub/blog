package web

import (
	"fmt"
	"github.com/MccGithub/blog/web/api"
	"github.com/go-chi/chi"
	"net/http"
)

type Opt struct {
	Address string
}

func Serve(opt Opt) error {
	handler := chi.NewRouter()

	handler.Mount("/api", api.Router())

	fmt.Println("Listening at ", opt.Address)
	return http.ListenAndServe(opt.Address, handler)
}