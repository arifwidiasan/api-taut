package repository

import (
	"fmt"

	"github.com/arifwidiasan/api-taut/model"
	"gorm.io/gorm/clause"
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

func (r *repositoryMysqlLayer) GetSosmedByUserID(user_id int) (sosmed model.Sosmed, err error) {
	res := r.DB.Where("user_id = ?", user_id).Preload(clause.Associations).Find(&sosmed)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("sosmed not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateSosmedByUserID(id int, sosmed model.Sosmed) error {
	res := r.DB.Where("user_id = ?", id).UpdateColumns(&sosmed)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update sosmed, no input or user id not found")
	}

	return nil
}
