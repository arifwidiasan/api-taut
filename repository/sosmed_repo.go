package repository

import (
	"fmt"

	"github.com/arifwidiasan/api-taut/model"
)

func (r *repositoryMysqlLayer) CreateSosmed(sosmed model.Sosmed) error {
	res := r.DB.Create(&sosmed)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error create sosmed")
	}

	return nil
}
