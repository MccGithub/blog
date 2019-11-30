package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"strings"
	"time"

	"github.com/chive-chan/blog/internal/dao"
	_ "github.com/mattn/go-sqlite3"
	"github.com/russross/blackfriday"
)

var (
	author string
	brief  string
	file   string
)

func init() {
	flag.StringVar(&author, "a", "韭菜", "listening address")
	flag.StringVar(&brief, "b", "", "文章简介")
	flag.StringVar(&file, "f", "", "md文件")

	flag.Parse()
}

func main() {
	db, err := dao.NewSQLHelper("sqlite3", "data.db")
	if err != nil {
		panic(err)
	}
	title := file[strings.LastIndex(file, "/")+1 : strings.LastIndex(file, ".")]
	article := db.GetArticle(author, title)
	var markdown []byte
	if markdown, err = ioutil.ReadFile(file); err != nil {
		panic(err)
	}
	article.Content = template.HTML(blackfriday.Run(markdown))
	article.Time = time.Now().Unix()
	article.Brief = brief
	if err = article.Insert(); err != nil {
		panic(err)
	}
}
