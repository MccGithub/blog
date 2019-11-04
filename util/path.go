package util

import (
	"html/template"
	"io"
	"path/filepath"
)

var (
	BasePath         string
	DbRelativePath string
	TmplRelativePath string
)

type Data map[string]interface{}

func GetDbPath(db string) string {
	return filepath.Join(BasePath, DbRelativePath, db)
}

func GetTmplPath(file string) string {
	return filepath.Join(BasePath, TmplRelativePath, file)
}

func View(w io.Writer, data Data, file ...string) error {
	for i, v := range file {
		file[i] = GetTmplPath(v)
	}
	tmpl, err := template.ParseFiles(file...)
	if err != nil {
		return err
	}

	return tmpl.Execute(w, data)
}
//模板嵌套时需要传入母子模板多个文件路径解析,故修改参数列表以支持传入多个文件路径参数
//func View(file string, w io.Writer, data Data) error {
//	tmpl, err := template.ParseFiles(GetTempPath(file))
//	if err != nil {
//		return err
//	}
//
//	return tmpl.Execute(w, data)
//}

// 改为数据库实现后不需要再使用
//func GetFileListByPath(path string) ([]os.FileInfo, error) {
//	files, err := ioutil.ReadDir(path)
//	return files, err
//}
