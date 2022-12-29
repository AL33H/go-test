package database

import (
	"fmt"

	"github.com/al33h/go-test/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetConnection() {
	dsn := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Banco conectado!")
	DB = db
	DB.AutoMigrate(&internal.ConsultaRealizada{})
}
