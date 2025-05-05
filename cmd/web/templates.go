package main

import (
	"text/template"
	"time"

	"github.com/yikuanzz/snippet/internal/models"
	"github.com/yikuanzz/snippet/ui"
)

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
	User            *models.User
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := ui.Templates.ReadDir("html/pages")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		if page.IsDir() {
			continue
		}
		name := page.Name()

		// 先用 template.New(name).Funcs(functions) 创建模板集合
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Templates, "html/base.tmpl")
		if err != nil {
			return nil, err
		}
		// 加载 partials
		ts, err = ts.ParseFS(ui.Templates, "html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}
		// 加载页面本身
		ts, err = ts.ParseFS(ui.Templates, "html/pages/"+name)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
