package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (c *Comment) Create() (err error) {
	if c.Post == nil {
		err = errors.New("no post found")
		return
	}
	err = Db.QueryRow(
		"INSERT INTO comments (content, author, post_id) VALUES ($1, $2, $3) RETURNING id",
		c.Content, c.Author, c.Post.Id).Scan(&c.Id)
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id = $1",
		id).Scan(&post.Id, &post.Content, &post.Author)
	if err != nil {
		return
	}

	rows, err := Db.Query("select id, content, author from comments where post_id = $1", post.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	_ = rows.Close()
	return
}

func (p *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id",
		p.Content, p.Author).Scan(&p.Id)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	_ = post.Create()

	comment := Comment{Content: "Good Post!", Author: "Joe", Post: &post}
	_ = comment.Create()

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
}
