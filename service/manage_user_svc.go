package service

import (
	"fmt"
	"strings"

	"github.com/arifwidiasan/api-taut/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *svc) AdminCreateUserService(user model.User) error {
	user.Username = strings.ToLower(user.Username)
	_, err := s.repo.GetAdminByUsername(user.Username)
	if err == nil {
		return fmt.Errorf("username already exist")
	}

	if user.Email == "" || user.Username == "" || user.Name == "" || user.Job == "" || user.PhoneNumber == "" {
		return fmt.Errorf("field username, email, name, job, and phone number cannot be empty")
	}

	user.Password = user.Username + "2023"
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error generate password")
	}
	user.Password = string(hash)

	return s.repo.CreateUser(user)
}

func (s *svc) AdminGetAllUserService() []model.User {
	return s.repo.GetAllUser()
}

func (s *svc) AdminGetUserByIDService(id int) (model.User, error) {
	return s.repo.GetUserByID(id)
}
func (s *svc) AdminUpdateUserByIDService(id int, user model.User) error {
	if user.Username != "" {
		return fmt.Errorf("username cannot be changed")
	}

	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("error generate password")
		}
		user.Password = string(hash)
	}

	return s.repo.UpdateUserByID(id, user)
}
