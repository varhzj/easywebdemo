package models

import (
	"time"

	"fmt"
	"net/url"
	"strings"

	"log"

	"github.com/varhzj/easywebdemo/conf"
)

type Post struct {
	Id         int
	UserId     int
	Author     string
	Title      string
	Color      string
	UrlName    string
	UrlType    int8
	Content    string
	Tags       string
	Views      int
	Status     int8
	PostTime   time.Time
	UpdateTime time.Time
	IsTop      int8
}

func (p *Post) ReadById() {
	db := conf.GetDBStore().GetDB()
	row := db.QueryRow("select title, content from t_post where id = ?", p.Id)
	row.Scan(&p.Title, &p.Content)
}

func (p *Post) ReadByPage(start, end int) []*Post {
	var (
		list []*Post
		err  error
	)
	db := conf.GetDBStore().GetDB()
	rows, err := db.Query("select id, title, content, post_time, tags from t_post where status = 0 and url_type = 0 order by id desc limit ?, ?", start, end)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var post = new(Post)
		rows.Scan(&post.Id, &post.Title, &post.Content, &post.PostTime, &post.Tags)
		list = append(list, post)
	}
	return list
}

func (p *Post) ReadByTag(tag string, start, end int) []*Post {
	var (
		list []*Post
		err  error
	)
	db := conf.GetDBStore().GetDB()
	rows, err := db.Query("select id, title, content, post_time, tags from t_post where status = 0 and url_type = 0 and tags like ? order by id limit ?, ?", "%"+tag+"%", start, end)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var post = new(Post)
		rows.Scan(&post.Id, &post.Title, &post.Content, &post.PostTime, &post.Tags)
		list = append(list, post)
	}
	return list
}

//带颜色的标题
func (m *Post) ColorTitle() string {
	if m.Color != "" {
		return fmt.Sprintf("<span style=\"color:%s\">%s</span>", m.Color, m.Title)
	} else {
		return m.Title
	}
}

//内容URL
func (m *Post) Link() string {
	if m.UrlName != "" {
		if m.UrlType == 1 {
			return fmt.Sprintf("/%s", strings.Replace(url.QueryEscape(m.UrlName), "+", "%20", -1))
		}
		return fmt.Sprintf("/article/%s", strings.Replace(url.QueryEscape(m.UrlName), "+", "%20", -1))
	}
	return fmt.Sprintf("/article/%d", m.Id)
}
