package core

import (
	"fmt"
	"html/template"
	"net/http"
)

// 渲染视图
func RenderHtml(w http.ResponseWriter, tpl string, data interface{}) {
	rtpl := fmt.Sprintf("templates/%s.html", tpl)

	tmpl, err := template.ParseFiles(TplPath, rtpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
