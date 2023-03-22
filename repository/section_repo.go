package repository

import (
	"fmt"

	"github.com/arifwidiasan/api-taut/model"
	"gorm.io/gorm/clause"
)

func (r *repositoryMysqlLayer) CreateSection(section model.Section) error {
	res := r.DB.Create(&section)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error create section")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAllSectionByUserID(id int) []model.Section {
	sections := []model.Section{}
	r.DB.Where("user_id = ?", id).Preload(clause.Associations).Find(&sections)

	return sections
}
