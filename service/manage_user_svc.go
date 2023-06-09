package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/arifwidiasan/api-taut/helper"
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

	user.QrcodePathFile, err = helper.GenerateQRCode(user.Username)
	if err != nil {
		return fmt.Errorf("error generate qrcode")
	}

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

func (s *svc) AdminDeleteUserByIDService(id int) error {
	_ = s.DeleteSosmedByUserIDService(id)
	_ = s.DeleteAllSectionByUserIDService(id)

	user, _ := s.repo.GetUserByID(id)
	_ = os.Remove("../uploads/qrcode/" + user.Username + ".png")
	_ = os.Remove("../uploads/profile-picture/" + user.Username + ".png")

	return s.repo.DeleteUserByID(id)
}
