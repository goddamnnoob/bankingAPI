package service

import "github.com/goddamnnoob/notReddit/domain"

type UserService interface {
	GetAllUsers() ([]domain.User, error)
	GetUser(string) (*domain.User, error)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (u DefaultUserService) GetAllUsers() ([]domain.User, error) {
	return u.repo.GetAllUsers()
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}

func (s DefaultUserService) GetUser(id string) (*domain.User, error) {
	return s.repo.ById(id)
}
