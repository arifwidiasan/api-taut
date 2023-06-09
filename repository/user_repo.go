package repository

import (
	"fmt"

	"github.com/arifwidiasan/api-taut/model"
)

func (r *repositoryMysqlLayer) CreateUser(user model.User) error {
	res := r.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error create user, maybe username or email already exist")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAllUser() []model.User {
	var users []model.User
	r.DB.Find(&users)

	return users
}

func (r *repositoryMysqlLayer) GetUserByID(id int) (user model.User, err error) {
	res := r.DB.Where("id = ?", id).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("user not found")
	}

	return
}

func (r *repositoryMysqlLayer) GetUserByUsername(username string) (user model.User, err error) {
	res := r.DB.Where("username = ?", username).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("user not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateUserByID(id int, user model.User) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update user, maybe updated username or email already exist")
	}

	return nil
}

func (r *repositoryMysqlLayer) DeleteUserByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.User{}, id)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete user, id user not found")
	}

	return nil
}
