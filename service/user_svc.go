package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/arifwidiasan/api-taut/helper"
	"github.com/arifwidiasan/api-taut/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *svc) CreateUserService(user model.User) error {
	user.Username = strings.ToLower(user.Username)
	_, err := s.repo.GetAdminByUsername(user.Username)
	if err == nil {
		return fmt.Errorf("username already exist")
	}

	if user.Email == "" || user.Username == "" || user.Name == "" || user.Job == "" || user.PhoneNumber == "" || user.Password == "" {
		return fmt.Errorf("field username, password, email, name, job, and phone number cannot be empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error generate password")
	}

	user.Password = string(hash)
	err = s.repo.CreateUser(user)
	if err != nil {
		return err
	}

	created_user, _ := s.repo.GetUserByUsername(user.Username)
	sosmed := model.Sosmed{}
	sosmed.UserID = created_user.ID
	_ = s.CreateSosmedService(sosmed)

	return nil
}

func (s *svc) LoginUserService(username, password string) (string, int) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", http.StatusUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateTokenUser(int(user.ID), user.Username, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}

func (s *svc) GetUserByUsernameService(username string) (model.User, error) {
	return s.repo.GetUserByUsername(username)
}

func (s *svc) ChangePassUserService(username, oldpass, newpass string) error {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldpass))
	if err != nil {
		return fmt.Errorf("old password not match")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(newpass), bcrypt.DefaultCost)
	user.Password = string(hash)

	err = s.repo.UpdateUserByID(int(user.ID), user)
	if err != nil {
		return fmt.Errorf("error update password user")
	}

	return nil
}

func (s *svc) UpdateUserByUsernameService(username string, user model.User) error {
	if user.Username != "" {
		return fmt.Errorf("username cannot be changed")
	}

	if user.Password != "" {
		return fmt.Errorf("password cannot be changed from here")
	}

	id, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	err = s.repo.UpdateUserByID(int(id.ID), user)
	if err != nil {
		return err
	}

	return nil
}
