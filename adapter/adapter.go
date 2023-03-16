package adapter

import (
	"github.com/arifwidiasan/api-taut/model"
	"github.com/golang-jwt/jwt"
)

type AdapterRepository interface {
	GetAdminByUsername(username string) (admin model.Admin, err error)
	UpdateAdminByID(id int, admin model.Admin) error
	CreateAdmin(admin model.Admin) error
	GetAllAdmin() []model.Admin
	GetAdminByID(id int) (admin model.Admin, err error)
	DeleteAdminByID(id int) error

	CreateUser(user model.User) error
	GetAllUser() []model.User
	GetUserByID(id int) (user model.User, err error)
	GetUserByUsername(username string) (user model.User, err error)
	UpdateUserByID(id int, user model.User) error
}

type AdapterService interface {
	ClaimToken(bearer *jwt.Token) string

	LoginAdmin(username, password string) (string, int)
	ChangePassAdminService(username, oldpass, newpass string) error
	CreateAdminService(admin model.Admin) error
	GetAdminByUsernameService(username string) (model.Admin, error)
	GetAllAdminService() []model.Admin
	GetAdminByIDService(id int) (model.Admin, error)
	UpdateAdminByIDService(id int, admin model.Admin) error
	DeleteAdminByIDService(id int) error

	AdminCreateUserService(user model.User) error
	AdminGetAllUserService() []model.User
	AdminGetUserByIDService(id int) (model.User, error)
	AdminUpdateUserByIDService(id int, user model.User) error
}
