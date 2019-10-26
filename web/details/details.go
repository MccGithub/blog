package details

import (
	"github.com/MccGithub/blog/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Details(w http.ResponseWriter, r *http.Request) {
	tmplFiles := []string{
		"html/index.html",
		"html/articles.html",
		"html/details.html",
	}
	// 传入模拟数据以测试, 第一次是空数据
	test_data := util.Data{
		"target": "details",
		"name": "文章详情页测试",
		"author": "作者当然是我啦",
		"content": `测试内容就是乱七八糟%>_<%<br />
jsaiofhhaoihfa<br />
叫啥佛我后爱疯hi哈搜发完后发动哈佛按时非 哈市豆腐<br />
if哈哈是佛我荷藕我更好歌手阿福<br />
后端和覅ua<br />`,
	}
	//if err := util.View(w, test_data, "html/index.html", "html/articles.html", "html/details.html"); err != nil {
	if err := util.View(w, test_data, tmplFiles...); err != nil {
		logrus.Warn(err)
	}

}
