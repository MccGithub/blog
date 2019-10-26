package util

import (
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	BasePath         string
	DbRelativePath string
	TmplRelativePath string
	ArtiRelativePath string
	tmplBasePath string
	artiBasePath string
	dbBasePath string
)

type Data map[string]interface{}

func GetTmplBasePath() string {
	if tmplBasePath == "" {
		tmplBasePath = filepath.Join(BasePath, TmplRelativePath)
	}
	return tmplBasePath
}

func GetDbBasePath() string {
	if dbBasePath == "" {
		dbBasePath = filepath.Join(BasePath, DbRelativePath)
	}
	return dbBasePath
}

func GetTmplPath(file string) string {
	return filepath.Join(BasePath, TmplRelativePath, file)
}

func GetArtiBasePath() string {
	if artiBasePath == "" {
		artiBasePath = filepath.Join(BasePath, ArtiRelativePath)
	}
	return artiBasePath
}

func GetArtiPath(file string) string {
	return filepath.Join(BasePath, ArtiRelativePath, file)
}

func GetFileListByPath(path string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	return files, err
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
