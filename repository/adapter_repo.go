package repository

import (
	"github.com/arifwidiasan/api-taut/adapter"

	"gorm.io/gorm"
)

type repositoryMysqlLayer struct {
	DB *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) adapter.AdapterRepository {
	return &repositoryMysqlLayer{
		DB: db,
	}
}
