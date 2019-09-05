package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"  // 前面带有横杠代表调用初始化
	"log"
)

var Db1 *sql.DB

func main()  {
	var err error
	Db1, err := sql.Open("mysql", "root:root(127.0.0.1:3306)/chitchat?parseTime=true")
	if err != nil{
		log.Fatal(err)
	}
	post :=Post{
		Content:"hello world",
		Author:"vanyarpy",
	}
	post.Create()
	defer Db1.Close()

}

type Post1 struct {
	Id int
	Content string
	Author string
}

func (post *Post1)Create() (err error)  {
	rs, err := Db.Exec("INSERT INTO posts (content, author) Values (?, ?)", post.Content, post.Author)
	if err != nil{
		log.Fatalln(err)
	}
	id, err := rs.LastInsertId()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(id)
	return
}

func insert()  {
	stmt, err := db.Prepare("insert into user(name,age)values(?,?)")
	if err != nil {
		log.Println(err)
	}

	rs, err := stmt.Exec("go-test", 12)
	if err != nil {
		log.Println(err)
	}
	//我们可以获得插入的id
	id, err := rs.LastInsertId()
	//可以获得影响行数
	affect, err := rs.RowsAffected()
}