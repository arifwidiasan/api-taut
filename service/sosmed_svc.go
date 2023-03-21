package service

import (
	"fmt"

	"github.com/arifwidiasan/api-taut/model"
)

func (s *svc) CreateSosmedService(sosmed model.Sosmed) error {
	return s.repo.CreateSosmed(sosmed)
}

func (s *svc) DeleteSosmedByUserIDService(user_id int) error {
	return s.repo.DeleteSosmedByUserID(user_id)
}

func (s *svc) GetSosmedByUserIDService(user_id int) (model.Sosmed, error) {
	return s.repo.GetSosmedByUserID(user_id)
}

func (s *svc) GetSosmedByUsernameService(username string) (model.Sosmed, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return model.Sosmed{}, err
	}

	return s.repo.GetSosmedByUserID(user.ID)
}

func (s *svc) UpdateSosmedByUsernameService(username string, sosmed model.Sosmed) error {
	if sosmed.UserID != 0 {
		return fmt.Errorf("input user id not allowed")
	}

	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	return s.repo.UpdateSosmedByUserID(user.ID, sosmed)
}
