package database

import (
    // "fmt"
    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
  err := godotenv.Load()
    USER     := "user"
    PASS     := "password"
    PROTOCOL := "tcp(mysql:3306)"
    DBNAME   := "database"

    dsn := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true"
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
  
    if err != nil {
      panic(err.Error())
    }
}