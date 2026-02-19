package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})

	// For later postgre / supabase
	// conn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable"
	// db, err := gorm.Open(postgre.Open(conn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database. Check the connection or the connection field")
	}

	// go feature for auto migration
	db.AutoMigrate(&Products{})
	return db
}
