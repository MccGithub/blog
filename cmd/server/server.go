package main

import (
	"flag"
	"os"
	"runtime"
	"strings"

	"github.com/MccGithub/blog/util"
	"github.com/MccGithub/blog/web"
	"github.com/sirupsen/logrus"
)

var (
	logLevel string
	opt      web.Opt
)

func init() {
	flag.StringVar(&opt.Address, "listen", ":8080", "listening address")
	flag.StringVar(&opt.DBDriver, "db.driver", "sqlite3", "database driver. supported: sqlit3")
	flag.StringVar(&opt.DBConn, "db.conn", "data.db", "connection string")
	flag.StringVar(&logLevel, "log.level", "warn", "logging level: trace, debug, info, error, fatal")

	flag.Parse()
}

func getCurrent() {
	// getCurrent函数用于获取定义此函数的文件的绝对路径
	// 故必须在main中定义并在main函数首行执行
	_, util.BasePath, _, _ = runtime.Caller(0)
	i := strings.LastIndex(util.BasePath, string(os.PathSeparator))
	util.BasePath = util.BasePath[0 : i+1]
}

func main() {
	getCurrent()
	// util包中根据BasePath和TmplRelativePath来确定模板文件夹绝对路径
	// util包中根据BasePath和DbRelativePath来确定模板文件夹绝对路径
	util.TmplRelativePath = "../../web/template"
	util.DbRelativePath = "../../"
	if lvl, err := logrus.ParseLevel(logLevel); err != nil {
		logrus.Fatal(err)
	} else {
		logrus.SetLevel(lvl)
	}
	if err := web.Serve(opt); err != nil {
		panic(err)
	}
}
