package service

import (
	"github.com/goddamnnoob/notReddit/domain"
	"github.com/goddamnnoob/notReddit/dto"
	"github.com/goddamnnoob/notReddit/errs"
)

type UserService interface {
	GetAllUsers() ([]domain.User, *errs.AppError)
	GetUser(string) (*dto.UserResponse, *errs.AppError)
	GetUserByStatus(int) ([]domain.User, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (u DefaultUserService) GetAllUsers() ([]domain.User, *errs.AppError) {
	return u.repo.GetAllUsers()
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}

func (s DefaultUserService) GetUser(id string) (*dto.UserResponse, *errs.AppError) {
	u, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := u.ToDto()
	return &response, nil
}

func (s DefaultUserService) GetUserByStatus(status int) ([]domain.User, *errs.AppError) {
	return s.repo.ByStatus(status)
}
