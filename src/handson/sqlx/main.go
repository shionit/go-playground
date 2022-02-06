package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	Id         int
	Content    string
	AuthorName string `db:"author"`
}

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRowx("select id, content, author from posts where id = $1", id).StructScan(&post)
	if err != nil {
		return
	}
	return
}

func (p *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id",
		p.Content, p.AuthorName).Scan(&p.Id)
	return
}

func main() {
	post := Post{Content: "Hello World!", AuthorName: "Sau Sheong"}
	_ = post.Create()
	fmt.Println(post)

	readPost := Post{Id: 1}
	readPost, _ = GetPost(post.Id)
	fmt.Println(readPost)
}
