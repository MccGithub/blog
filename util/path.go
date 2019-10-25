package util

import (
	"html/template"
	"io"
	"path/filepath"
)

var (
	TmplBasePath string
	TmplRelativePath string
)

func GetTempPath(file string) string {
	return filepath.Join(TmplBasePath, TmplRelativePath, file)
}

func View(file string, w io.Writer, data interface{}) error {
	tmpl, err := template.ParseFiles(GetTempPath(file))
	if err != nil {
		return err
	}

	return tmpl.Execute(w, data)
}
