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

func (r *repositoryMysqlLayer) DeleteSosmedByUserID(user_id int) error {
	res := r.DB.Unscoped().Where("user_id = ?", user_id).Delete(&model.Sosmed{})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete sosmed, user_id not found")
	}

	return nil
}
