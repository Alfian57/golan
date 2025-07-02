package service

import (
	"context"

	"github.com/Alfian57/belajar-golang/internal/dto"
	"github.com/Alfian57/belajar-golang/internal/model"
	"github.com/Alfian57/belajar-golang/internal/repository"
	"github.com/Alfian57/belajar-golang/internal/utils"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{
		repository: r,
	}
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	users, err := s.repository.GetAll(ctx)
	return users, err
}

func (s *UserService) CreateUser(ctx context.Context, request dto.CreateUserRequest) error {
	hashedPass, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}

	user := model.User{
		Username: request.Username,
		Password: hashedPass,
	}

	err = s.repository.Create(ctx, &user)
	return err
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (model.User, error) {
	user, err := s.repository.GetByID(ctx, id)
	return user, err
}

func (s *UserService) UpdateUser(ctx context.Context, request dto.UpdateUserRequest) error {
	user := model.User{
		ID:       request.ID,
		Username: request.Username,
	}

	err := s.repository.Update(ctx, &user)
	return err
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	return err
}
