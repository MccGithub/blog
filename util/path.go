package util

import (
	"html/template"
	"io"
	"path/filepath"
)

var (
	BasePath         string
	TmplRelativePath string
	ArtiRelativePath string
)

type Data map[string]interface{}

func GetTempPath(file string) string {
	return filepath.Join(BasePath, TmplRelativePath, file)
}

func GetArtiPath(file string) string {
	return filepath.Join(BasePath, ArtiRelativePath, file)
}

func View(w io.Writer, data Data, file ...string) error {
	for i, v := range file {
		file[i] = GetTempPath(v)
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
