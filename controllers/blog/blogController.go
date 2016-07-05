package blog

import (
	"net/http"
	"strconv"

	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/varhzj/easywebdemo/models"
	"github.com/varhzj/easywebdemo/templates"
)

func ReadArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v := vars["id"]
	id, err := strconv.Atoi(v)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
	}
	post := &models.Post{Id: id}
	post.ReadById()
	templates.RenderTemplate(w, "article", "layout", post)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	post := &models.Post{}
	list := post.ReadByPage(0, -1)
	templates.RenderTemplate(w, "index", "layout", list)
}
