package repository

import (
	"fmt"

	"github.com/arifwidiasan/api-taut/model"
)

func (r *repositoryMysqlLayer) CreateSection(section model.Section) error {
	res := r.DB.Create(&section)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error create section")
	}

	return nil
}
