package service

import (
	"context"
	"time"

	"github.com/Alfian57/belajar-golang/internal/dto"
	errs "github.com/Alfian57/belajar-golang/internal/errors"
	"github.com/Alfian57/belajar-golang/internal/logger"
	"github.com/Alfian57/belajar-golang/internal/model"
	"github.com/Alfian57/belajar-golang/internal/repository"
)

type TodoService struct {
	todoRepository *repository.TodoRepository
	userRepository *repository.UserRepository
}

func NewTodoService(todoRepository *repository.TodoRepository, userRepository *repository.UserRepository) *TodoService {
	return &TodoService{
		todoRepository: todoRepository,
		userRepository: userRepository,
	}
}

func (s *TodoService) GetAllTodos(ctx context.Context) ([]model.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	todos, err := s.todoRepository.GetAll(ctx)

	if err != nil {
		logger.Log.Errorw("failed to get all todos", "error", err)
		return nil, errs.NewAppError(500, "failed to retrieve todos", err)
	}

	return todos, nil
}

func (s *TodoService) CreateTodo(ctx context.Context, request dto.CreateTodoRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := s.userRepository.GetByID(ctx, request.UserID)
	if err != nil && err != errs.ErrUserNotFound {
		fieldError := errs.NewFieldError("user_id", "user_id not exist")
		return errs.NewValidationError([]errs.FieldError{fieldError})
	}

	todo := model.Todo{
		Todo:   request.Todo,
		UserID: request.UserID,
	}

	if err := s.todoRepository.Create(ctx, &todo); err != nil {
		logger.Log.Errorw("failed to create todo", "todo", request.Todo, "user_id", request.UserID, "error", err)
		return errs.NewAppError(500, "failed to create todo", err)
	}

	logger.Log.Infow("todo created successfully", "todo", request.Todo)
	return nil
}

func (s *TodoService) GetTodoByID(ctx context.Context, id string) (model.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := s.todoRepository.GetByID(ctx, id)
	if err != nil {
		if err == errs.ErrTodoNotFound {
			return model.Todo{}, err
		}
		logger.Log.Errorw("failed to get todo by ID", "id", id, "error", err)
		return model.Todo{}, errs.NewAppError(500, "failed to retrieve todo", err)
	}
	return user, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, request dto.UpdateTodoRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := s.todoRepository.GetByID(ctx, request.ID.String())
	if err != nil {
		if err == errs.ErrTodoNotFound {
			return err
		}
		logger.Log.Errorw("failed to check todo existence for update", "id", request.ID, "error", err)
		return errs.NewAppError(500, "failed to validate todo", err)
	}

	todo := model.Todo{
		ID:     request.ID,
		Todo:   request.Todo,
		UserID: request.UserID,
	}

	if err := s.todoRepository.Update(ctx, &todo); err != nil {
		logger.Log.Errorw("failed to update todo", "id", request.ID, "error", err)
		return errs.NewAppError(500, "failed to update todo", err)
	}

	logger.Log.Infow("todo updated successfully", "id", request.ID)
	return nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := s.todoRepository.Delete(ctx, id); err != nil {
		if err == errs.ErrTodoNotFound {
			return err
		}
		logger.Log.Errorw("failed to delete todo", "id", id, "error", err)
		return errs.NewAppError(500, "failed to delete todo", err)
	}

	logger.Log.Infow("todo deleted successfully", "id", id)
	return nil
}
