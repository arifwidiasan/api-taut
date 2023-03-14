package service

import (
	"fmt"
	"net/http"

	"github.com/arifwidiasan/api-taut/helper"
	"github.com/arifwidiasan/api-taut/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *svc) LoginAdmin(username, password string) (string, int) {
	admin, err := s.repo.GetAdminByUsername(username)
	if err != nil {
		return "", http.StatusUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateTokenAdmin(int(admin.ID), admin.Username, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}

func (s *svc) ChangePassAdminService(username, oldpass, newpass string) error {
	admin, err := s.repo.GetAdminByUsername(username)
	if err != nil {
		return fmt.Errorf("admin not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(oldpass))
	if err != nil {
		return fmt.Errorf("old password not match")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(newpass), bcrypt.DefaultCost)
	admin.Password = string(hash)

	err = s.repo.UpdateAdminByID(int(admin.ID), admin)
	if err != nil {
		return fmt.Errorf("error update password admin")
	}

	return nil
}

func (s *svc) CreateAdminService(admin model.Admin) error {
	if admin.Username == "" || admin.Password == "" {
		return fmt.Errorf("username or password is empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error generate password")
	}

	admin.Password = string(hash)

	return s.repo.CreateAdmin(admin)
}

func (s *svc) GetAdminByUsernammeService(username string) (model.Admin, error) {
	return s.repo.GetAdminByUsername(username)
}
