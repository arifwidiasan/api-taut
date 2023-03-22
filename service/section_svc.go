package service

import (
	"fmt"

	"github.com/arifwidiasan/api-taut/model"
)

func (s *svc) CreateSectionService(username string, section model.Section) error {
	user, err := s.GetUserByUsernameService(username)
	if err != nil {
		return err
	}

	if section.Title == "" || section.Description == "" {
		return fmt.Errorf("title and description cannot be empty")
	}

	section.UserID = user.ID

	return s.repo.CreateSection(section)
}

func (s *svc) GetAllSectionByUserIDService(username string) []model.Section {
	user, err := s.GetUserByUsernameService(username)
	if err != nil {
		return nil
	}

	return s.repo.GetAllSectionByUserID(user.ID)
}
