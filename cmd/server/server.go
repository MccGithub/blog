package main

import (
	"flag"
	"github.com/MccGithub/blog/web"
	"github.com/sirupsen/logrus"
)

var (
	opt web.Opt
)

func init() {
	flag.StringVar(&opt.Address, "listen", ":8080", "listening address")

	flag.Parse()
}

func main() {
	if err := web.Serve(opt); err != nil {
		panic(err)
	}
}