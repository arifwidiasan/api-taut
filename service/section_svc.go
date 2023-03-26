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

func (s *svc) GetOneSectionByUserIDandIDService(username string, id int) (model.Section, error) {
	user, err := s.GetUserByUsernameService(username)
	if err != nil {
		return model.Section{}, err
	}

	return s.repo.GetOneSectionByUserIDandID(id, user.ID)
}

func (s *svc) UpdateSectionByUserIDandIDService(username string, id int, section model.Section) error {
	if section.UserID != 0 {
		return fmt.Errorf("input user id not allowed")
	}

	user, err := s.GetUserByUsernameService(username)
	if err != nil {
		return err
	}

	return s.repo.UpdateSectionByUserIDandID(id, user.ID, section)
}

func (s *svc) DeleteSectionByUserIDandIDService(username string, id int) error {
	user, err := s.GetUserByUsernameService(username)
	if err != nil {
		return err
	}

	return s.repo.DeleteSectionByUserIDandID(id, user.ID)
}
