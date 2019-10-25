package page

import (
	"github.com/MccGithub/blog/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Article(w http.ResponseWriter, r *http.Request) {
	// 获取存放文章文件夹中的文件信息
	filesInfo, err := util.GetFileListByPath(util.GetArtiBasePath())
	if err != nil {
		logrus.Warn(err)
	}

	// 提取出文件名存放在filesNames中
	var filesNames []string
	for _, v := range filesInfo {
		if !v.IsDir() {
			filesNames = append(filesNames, v.Name())
		}
	}

	// 初始化模板变量, 设置文件名列表
	data := util.Data{
		"files": filesNames,
	}

	// 读取url中的 file 参数值
	fileName := r.FormValue("file")
	if fileName == "" {
		fileName = "welcome.txt"
	}
	data["file"] = fileName

	if err := util.View(w, data, "html/main.html", "html/article.html"); err != nil {
		logrus.Warn(err)
	}
}