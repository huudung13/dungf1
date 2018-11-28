package helper

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type (
	TmplHelper struct {
		tmpl   *template.Template
		config TmplConfig
	}
	TmplConfig struct {
		Dir      string
		Suffix   string
		NotFound string
	}
)

func NewTPL(config TmplConfig, funcMap template.FuncMap) (tmpl *TmplHelper, err error) {
	tmpl = &TmplHelper{}
	validFiles := []string{}
	err = filepath.Walk(config.Dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, "."+config.Suffix) {
			validFiles = append(validFiles, path)
		}
		return nil
	})
	if err != nil {
		return
	}
	tmpl.tmpl = template.Must(template.New("tmpl").Funcs(funcMap).ParseFiles(validFiles...))
	tmpl.config = config
	return
}

func (t TmplHelper) Render(wr io.Writer, name string, data interface{}) {
	err := t.tmpl.ExecuteTemplate(wr, name, data)
	if err != nil {
		fmt.Println(err)
	}
}

func (t TmplHelper) NotFound(wr io.Writer) {
	t.Render(wr, t.config.NotFound, nil)
}
