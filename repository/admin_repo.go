package repository

import (
	"fmt"

	"github.com/arifwidiasan/api-taut/model"
)

func (r *repositoryMysqlLayer) GetAdminByUsername(username string) (admin model.Admin, err error) {
	res := r.DB.Where("username = ?", username).Find(&admin)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("admin not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateAdminByID(id int, admin model.Admin) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&admin)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update admin")
	}

	return nil
}

func (r *repositoryMysqlLayer) CreateAdmin(admin model.Admin) error {
	res := r.DB.Create(&admin)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error create admin, maybe username already exist")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAllAdmin() []model.Admin {
	admins := []model.Admin{}
	r.DB.Find(&admins)

	return admins
}
