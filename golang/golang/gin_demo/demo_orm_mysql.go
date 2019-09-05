package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func (post *Post) Create() (err error) {
	rs, err := Db.Exec("INSERT INTO posts (content, author) Values (?, ?)", post.Content, post.Author)
	if err != nil {
		log.Fatalln(err)
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id)
	return
}

var Db *sql.DB

type Title struct {
	gorm.Model
	id int
	title string
}

func main() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	// Db, err = sql.Open("mysql", "host=127.0.0.1 user=root dbname=title port=3306 sslmode=disable password=root")
	if err != nil {
		log.Fatal(err)
	}
	err = Db.Ping()
	if err != nil{
		log.Println(err)
	}
	defer Db.Close()

	Db.Au

	//post := Post{
	//	Content:"hello world",
	//	Author:"vanyarpy",
	//}
	//post.Create()

}