package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := DB.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	_ = rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = DB.QueryRow("select id, content, author from posts where id = $1", id).
		Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (p *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			return
		}
	}()
	err = stmt.QueryRow(p.Content, p.Author).Scan(&p.Id)
	return
}

func (p *Post) Update() (err error) {
	_, err = DB.Exec("update posts set content = $2, author = $3 where id = $1",
		p.Id, p.Content, p.Author)
	return
}

func (p *Post) Delete() (err error) {
	_, err = DB.Exec("delete from posts where id = $1", p.Id)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}

	fmt.Println(post)
	_ = post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	_ = readPost.Update()

	posts, _ := Posts(4)
	fmt.Println(posts)

	_ = readPost.Delete()
}
