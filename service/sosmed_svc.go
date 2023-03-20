package service

import (
	"github.com/arifwidiasan/api-taut/model"
)

func (s *svc) CreateSosmedService(sosmed model.Sosmed) error {
	return s.repo.CreateSosmed(sosmed)
}
