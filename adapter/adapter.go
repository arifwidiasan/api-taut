package adapter

import (
	"github.com/arifwidiasan/api-taut/model"
	"github.com/golang-jwt/jwt"
)

type AdapterRepository interface {
	GetAdminByUsername(username string) (admin model.Admin, err error)
	UpdateAdminByID(id int, admin model.Admin) error
}

type AdapterService interface {
	ClaimToken(bearer *jwt.Token) string

	LoginAdmin(username, password string) (string, int)
	ChangePassAdminService(username, oldpass, newpass string) error
}
