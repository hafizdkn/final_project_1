package database

import "gorm.io/gorm"

type Todos struct {
	gorm.Model
	Id        int
	Task      string
	Completed bool
}

func Generate(db *gorm.DB) {
	db.AutoMigrate(&Todos{})
}
