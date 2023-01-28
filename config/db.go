package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

  host := os.Getenv("DB_HOST")
  name := os.Getenv("DB_NAME")
  user := os.Getenv("DB_USER")
  pwd := os.Getenv("DB_PASSWORD")
  port := os.Getenv("DB_PORT")
  timezone := os.Getenv("DB_TIMEZONE")

  dsn := "host="+ host+ " user=" + user + " password=" + pwd + " dbname=" + name + " port=" + port + " sslmode=require TimeZone=" + timezone
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect to db")
  }

  DB = db
}
