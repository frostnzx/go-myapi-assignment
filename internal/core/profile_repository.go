package core

type ProfileRepository interface {
	Save(profile Profile) error 
	GetAll() ([]Profile, error)
}