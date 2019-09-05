package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Article struct {
	gorm.Model
	id int
	Article string
}

type Title1 struct {
	gorm.Model
	id int
	title string
}

type User struct {
	gorm.Model
	id int
	name string
}

type Role struct {
	gorm.Model
	id int
	role string
}

func (Role) TableName() string  {
	return "role"
}

func main()  {
	db, err := gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")
	if err != nil{
		fmt.Println(err)
	}
	defer db.Close()
	//db.AutoMigrate(&Article{})  // 数据库迁移
	//db.AutoMigrate(&Article{}, &Title1{})
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(Title1{})
	//// 创建
	// db.Create(&Article{id:12, Article:"这是一个好东西"})
	//// 查询
	// var article Article
	/*db.First(&article, 1)   // id 为1的文章
	fmt.Println(&article)*/
	//// 更新
	// db.Model(&article).Update("article", "这真是个好东西")
	//// 删除
	// db.Delete(&article)


	// 检查表是否存在
	//rls :=db.HasTable(&Article{})
	//rls :=db.HasTable("article")
	//if rls {
	//	fmt.Println("Article存在")
	//}else {
	//	fmt.Println("Article不存在")
	//}
	// 删除表
	//db.DropTable(&Title1{})
	//db.DropTable("Title")
	// db.DropTableIfExists(&Title1{}, &Article{})

	// 创建表
	db.CreateTable(&Role{})
	//db.CreateTable(&Title1{}, &Article{})
	//// 删除表列
	//db.Model(&Article{}).DropColumn("article")
	//
	//// 添加外键
	////db.Model(&Article{}).AddForeignKey("city_id", "cityes(id)")
	//db.Model(&Article{}).AddIndex("idx_user_name", "name")
	//// 为`name`, `age`列添加索引`idx_user_name_age`
	//db.Model(&User{}).AddIndex("idx_user_name_age", "name", "age")
	//// 添加唯一索引
	//db.Model(&User{}).AddUniqueIndex("idx_user_name", "name")
	//// 为多列添加唯一索引
	//db.Model(&User{}).AddUniqueIndex("idx_user_name_age", "name", "age")
	//// 删除索引
	//db.Model(&User{}).RemoveIndex("idx_user_name")

	//更改默认表明
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "prefix_" + defaultTableName;
	}
	// 重设列名
	type Animal struct {
		AnimalId    int64     `gorm:"column:beast_id"`         // 设置列名为`beast_id`
		Birthday    time.Time `gorm:"column:day_of_the_beast"` // 设置列名为`day_of_the_beast`
		Age         int64     `gorm:"column:age_of_the_beast"` // 设置列名为`age_of_the_beast`
	}
	// 给字段ID设置主键
	// 使用tag`primary_key`用来设置主键
	type Animal1 struct {
		AnimalId int64 `gorm:"primary_key"` // 设置AnimalId为主键
		Name     string
		Age      int64
	}

}
/*
表示结构体名称的复数形式
*/
// 模型类的定义
type User1 struct {
	gorm.Model
	Birthday time.Time
	Age int
	Name string  `gorm:"size:255"`
	Num int `gorm:"AUTO_INCREMENT"`

	CreditCard   CreditCard  // One-To-One (拥有一个 - CreditCard表的UserID作外键)
	Emails            []Email         // One-To-Many (拥有多个 - Email表的UserID作外键)

	BillingAddress    Address         // One-To-One (属于 - 本表的BillingAddressID作外键)
	BillingAddressID  sql.NullInt64

	ShippingAddress   Address         // One-To-One (属于 - 本表的ShippingAddressID作外键)
	ShippingAddressID int

	IgnoreMe          int `gorm:"-"`   // 忽略这个字段
	Languages         []Language `gorm:"many2many:user_languages;"` // Many-To-Many , 'user_languages'是连接表
}

type Email struct {
	ID      int
	UserID  int     `gorm:"index"` // 外键 (属于), tag `index`是为该列创建索引
	Email   string  `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` // 设置字段为非空并唯一
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` // 创建索引并命名，如果找到其他相同名称的索引则创建组合索引
	Code string `gorm:"index:idx_name_code"` // `unique_index` also works
}

type CreditCard struct {
	gorm.Model
	UserID  uint
	Number  string
}