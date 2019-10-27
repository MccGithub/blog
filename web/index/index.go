package index

import (
	"github.com/MccGithub/blog/internal/dao"
	"github.com/MccGithub/blog/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// 这是文章存储为txt文件时的实现
//func Articles(w http.ResponseWriter, r *http.Request) {
//	// 获取存放文章文件夹中的文件信息
//	filesInfo, err := util.GetFileListByPath(util.GetArtiBasePath())
//	if err != nil {
//		logrus.Warn(err)
//	}
//
//	// 提取出文件名(去掉后缀名后)存放在filesNames中
//	var filesNames []string
//	for _, v := range filesInfo {
//		if !v.IsDir() {
//			filesNames = append(filesNames, v.Name()[0:strings.LastIndex(v.Name(), ".")])
//		}
//	}
//
//	// 初始化模板变量, 设置文件名列表
//	data := util.Data{
//		"files": filesNames,
//	}
//
//	// 读取url中的 file 参数值
//	fileName := r.FormValue("file")
//	if fileName == "" {
//		fileName = "welcome"
//	}
//	data["file"] = fileName
//	content, err := ioutil.ReadFile(util.GetArtiPath(data["file"].(string))+".txt")
//	data["content"] = string(content)
//
//	if err := util.View(w, data, "html/index.html", "html/details.html"); err != nil {
//		logrus.Warn(err)
//	}
//}
func Articles(w http.ResponseWriter, r *http.Request) {
	tmplFiles := []string{
		"html/index.html",
		"html/articles.html",
		"html/details.html",
	}
	var articles []dao.Article
	db := r.Context().Value("db").(*dao.SQLHelper)
	for i := 0; i < 100; i++ {
		id := strconv.Itoa(i)
		article := db.GetArticle(id)
		article.Name = "name" + id
		article.Author = "author" + id
		article.Brief = "brief" + id
		article.Content = "content" + id
		articles = append(articles, *article)
	}
	// 传入模拟数据以测试
	test_data := util.Data{
		"target": "index",
		"articles": articles,
	}
	//if err := util.View(w, test_data, "html/index.html", "html/articles.html", "html/details.html"); err != nil {
	if err := util.View(w, test_data, tmplFiles...); err != nil {
		logrus.Warn(err)
	}
}
