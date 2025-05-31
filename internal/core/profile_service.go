package core

import "errors"

type ProfileService interface {
	CreateProfile(profile Profile) error // PUT new profile
	GetProfiles() ([]Profile, error) // GET all profiles
}

type ProfileServiceImpl struct {
	repo ProfileRepository
}
func NewProfileService(repo ProfileRepository) ProfileService {
	return &ProfileServiceImpl{
		repo: repo,
	}
}
func (s *ProfileServiceImpl) CreateProfile(order Profile) error {
	// validate
	if order.Name == "" {
		return errors.New("name must not be empty")
	}
	// save on db
	if err := s.repo.Save(order); err != nil {
		return err
	}
	return nil
}
func (s *ProfileServiceImpl) GetProfiles() ([]Profile, error) {
	profiles, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return profiles, nil
}