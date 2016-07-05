package templates

import (
	"html/template"
	"log"
	"net/http"
)

var (
	templates      map[string]*template.Template
	blogTplFuncMap = make(template.FuncMap)
)

func init() {

	blogTplFuncMap["str2html"] = Str2html
	blogTplFuncMap["date"] = Date
	log.Println(blogTplFuncMap)
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["index"] = template.Must(template.New("index").Funcs(blogTplFuncMap).ParseFiles("views/default/index.html", "views/default/layout.html"))
	templates["article"] = template.Must(template.New("article").Funcs(blogTplFuncMap).ParseFiles("views/default/article.html", "views/default/layout.html"))
}

func RenderTemplate(w http.ResponseWriter, name, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
