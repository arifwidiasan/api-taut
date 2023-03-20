package service

import (
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
