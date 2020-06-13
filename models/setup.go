package models


import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
  db, err := gorm.Open("mysql", "root@(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local")

  if  err != nil {
    panic("Failed to connect to DB")
  }

  db.AutoMigrate(&Book{})

  DB = db

}
