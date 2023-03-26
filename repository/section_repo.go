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

func (r *repositoryMysqlLayer) GetOneSectionByUserIDandID(id, user_id int) (section model.Section, err error) {
	res := r.DB.Where("id = ? AND user_id = ?", id, user_id).Preload(clause.Associations).Find(&section)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("section not found or dont have access to this section")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateSectionByUserIDandID(id int, user_id int, section model.Section) error {
	res := r.DB.Where("id = ? AND user_id = ?", id, user_id).UpdateColumns(section)
	if res.RowsAffected < 1 {
		return fmt.Errorf("section not found or dont have access to this section or no change")
	}

	return nil
}

func (r *repositoryMysqlLayer) DeleteSectionByUserIDandID(id, user_id int) error {
	res := r.DB.Unscoped().Where("id = ? AND user_id = ?", id, user_id).Delete(&model.Section{})
	if res.RowsAffected < 1 {
		return fmt.Errorf("section not found or dont have access to this section")
	}

	return nil
}

func (r *repositoryMysqlLayer) DeleteAllSectionByUserID(id int) error {
	res := r.DB.Unscoped().Where("user_id = ?", id).Delete(&model.Section{})
	if res.RowsAffected < 1 {
		return fmt.Errorf("section not found or dont have access to this section")
	}

	return nil
}
