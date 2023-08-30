package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student
}
type Student struct {
	gorm.Model
	StudentName string
	ClassID     uint
}
type IDCard struct {
	gorm.Model
	IDCardName string
}
type Teacher struct {
	gorm.Model
	TeacherName string
}

func main() {

	dsn := "root:111111@tcp(127.0.0.1:3306)/mygorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Class{}, &Student{}, &IDCard{}, &Teacher{})
	db.Create(&Class{
		ClassName: "jinzhu",
		Students:  []Student{{StudentName: "学生王五", ClassID: 28}},
	})

}
