package main

import (
	"github.com/jinzhu/gorm"
	"time"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"
import "fmt"

type User2 struct {
	gorm.Model
	Id int
	Name string
	Age int
	Brithday time.Time
}

type Animal struct {
	Id int64
	Name string `gorm:"default:'galeone'"`
	Age int64
}

var animal = Animal{Age:99, Name:""}

func (user *User) BrforeCreate(scope *gorm.Scope) error  {
	scope.SetColumn("ID", uuid.New())
	return nil
}

func main()  {
	db, err := gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")
	if err != nil{
		fmt.Println(err)
	}
	defer db.Close()
	db.CreateTable(&User2{})
	user := User2{Name:"Jinzhu", Age:18, Brithday: time.Now()}
	db.NewRecord(user)
	db.Create(&user)
	db.NewRecord(user)

	db.Create(&animal)

	// 查询操作，获取第一条记录
	db.First(&user)
	// 获取最后一条记录
	db.Last(&user)
	// 获取所有记录
	db.Find(&user)
	// 使用主键获取记录
	db.First(&user, 10)

	// where查询条件语句
	// 获取第一个匹配的结果
	db.Where("name = ?", "jinzhu").First(&user)
	// 获取所有匹配的结果
	db.Where("name = ?", "jinzhu").Find(&user)
	db.Where("name <> ?", "jinzhu").Find(&user)
	db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&user)
	db.Where("name LIKE ?", "%jin%").Find(&user)
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&user)
	db.Where("updated_at > ?", lastWeek).Find(&user)
	db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)

}

