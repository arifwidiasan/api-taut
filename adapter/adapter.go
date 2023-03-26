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
	DeleteUserByID(id int) error

	CreateSosmed(sosmed model.Sosmed) error
	DeleteSosmedByUserID(user_id int) error
	GetSosmedByUserID(user_id int) (sosmed model.Sosmed, err error)
	UpdateSosmedByUserID(id int, sosmed model.Sosmed) error

	CreateSection(section model.Section) error
	GetAllSectionByUserID(id int) []model.Section
	GetOneSectionByUserIDandID(id, user_id int) (section model.Section, err error)
	UpdateSectionByUserIDandID(id int, user_id int, section model.Section) error
	DeleteSectionByUserIDandID(id int, user_id int) error
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
	AdminDeleteUserByIDService(id int) error

	CreateSosmedService(sosmed model.Sosmed) error
	DeleteSosmedByUserIDService(user_id int) error
	GetSosmedByUserIDService(user_id int) (model.Sosmed, error)
	GetSosmedByUsernameService(username string) (model.Sosmed, error)
	UpdateSosmedByUsernameService(username string, sosmed model.Sosmed) error

	CreateUserService(user model.User) error
	LoginUserService(username, password string) (string, int)
	GetUserByUsernameService(username string) (model.User, error)
	ChangePassUserService(username, oldpass, newpass string) error
	UpdateUserByUsernameService(username string, user model.User) error

	CreateSectionService(username string, section model.Section) error
	GetAllSectionByUserIDService(username string) []model.Section
	GetOneSectionByUserIDandIDService(username string, id int) (model.Section, error)
	UpdateSectionByUserIDandIDService(username string, id int, section model.Section) error
	DeleteSectionByUserIDandIDService(username string, id int) error
}
