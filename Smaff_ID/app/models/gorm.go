package models

import (
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
)

var (
  DB *gorm.DB
)

func InitDB() {
  USER := "a.satonaka"
  PASSWORD := "satochan"
  HOSTNAME := "127.0.0.1"
  DBNAME := "Smaff_ID"

  dsn := USER + ":" + PASSWORD + "@" + HOSTNAME + "/" + DBNAME + "?parseTime=true&loc=Asia%2FTokyo"
  var err error

  DB,err = gorm.Open("mysql",dsn)
  //DB接続時にエラーがあった場合パニックを起こす
  if err != nil {
    panic(err)
  }
  //DB.LogMode(true)
}