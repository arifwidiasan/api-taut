package adapter

import "github.com/arifwidiasan/api-taut/model"

type AdapterRepository interface {
	GetAdminByUsername(username string) (admin model.Admin, err error)
	UpdateAdminByID(id int, admin model.Admin) error
}

type AdapterService interface {
	LoginAdmin(username, password string) (string, int)
	ChangePassAdminService(oldpass, newpass string) error
}
