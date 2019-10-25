package main

import (
	"flag"
	"github.com/MccGithub/blog/util"
	"github.com/MccGithub/blog/web"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
)

var (
	logLevel string
	opt web.Opt
)

func init() {
	flag.StringVar(&opt.Address, "listen", ":80", "listening address")
	flag.StringVar(&logLevel, "log.level", "warn", "logging level: trace, debug, info, error, fatal")

	flag.Parse()
}

func getCurrent() {
	_, util.TmplBasePath, _, _ = runtime.Caller(0)
	i := strings.LastIndex(util.TmplBasePath, string(os.PathSeparator))
	util.TmplBasePath = util.TmplBasePath[0:i+1]
}

func main() {
	getCurrent()
	if lvl, err := logrus.ParseLevel(logLevel); err != nil {
		logrus.Fatal(err)
	} else {
		logrus.SetLevel(lvl)
	}
	if err := web.Serve(opt); err != nil {
		panic(err)
	}
}