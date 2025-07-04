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

func (s *TodoService) GetAllTodos(ctx context.Context, currentUser model.User, query dto.GetTodosFilter) (dto.PaginatedResult[model.Todo], error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query.PaginationRequest.SetDefaults()
	query.SetDefaults()

	todos, err := s.todoRepository.GetAllByUser(ctx, currentUser.ID.String(), query)
	if err != nil {
		logger.Log.Errorw("failed to get todos", "error", err)
		return dto.PaginatedResult[model.Todo]{}, errs.NewAppError(500, "failed to retrieve todos", err)
	}

	count, err := s.todoRepository.CountAllByUser(ctx, currentUser.ID.String(), query)
	if err != nil {
		logger.Log.Errorw("failed to count todos", "error", err)
		return dto.PaginatedResult[model.Todo]{}, errs.NewAppError(500, "failed to retrieve todos", err)
	}

	pagination := dto.NewPaginationResponse(query.Page, query.Limit, count)
	result := dto.PaginatedResult[model.Todo]{
		Data:       todos,
		Pagination: pagination,
	}

	return result, nil
}

func (s *TodoService) CreateTodo(ctx context.Context, currentUser model.User, request dto.CreateTodoRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	todo := model.Todo{
		Todo:   request.Todo,
		UserID: currentUser.ID,
	}

	if err := s.todoRepository.Create(ctx, &todo); err != nil {
		logger.Log.Errorw("failed to create todo", "todo", request.Todo, "user_id", currentUser.ID, "error", err)
		return errs.NewAppError(500, "failed to create todo", err)
	}

	logger.Log.Infow("todo created successfully", "todo", request.Todo)
	return nil
}

func (s *TodoService) GetTodoByID(ctx context.Context, currentUser model.User, id string) (model.Todo, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	todo, err := s.todoRepository.GetByID(ctx, id)
	if err != nil {
		if err == errs.ErrTodoNotFound {
			return model.Todo{}, err
		}
		logger.Log.Errorw("failed to get todo by ID", "id", id, "error", err)
		return model.Todo{}, errs.NewAppError(500, "failed to get todo", err)
	}

	if todo.UserID != currentUser.ID {
		logger.Log.Errorw("forbidden", "id", id, "error", errs.ErrForbidden)
		return model.Todo{}, errs.ErrForbidden
	}

	return todo, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, currentUser model.User, request dto.UpdateTodoRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	todo, err := s.todoRepository.GetByID(ctx, request.ID.String())
	if err != nil {
		if err == errs.ErrTodoNotFound {
			return err
		}
		logger.Log.Errorw("failed to get todo", "id", request.ID.String(), "error", err)
		return errs.NewAppError(500, "failed to get todo", err)
	}

	if todo.UserID != currentUser.ID {
		logger.Log.Errorw("forbidden", "id", request.ID.String(), "error", errs.ErrForbidden)
		return errs.ErrForbidden
	}

	_, err = s.todoRepository.GetByID(ctx, request.ID.String())
	if err != nil {
		if err == errs.ErrTodoNotFound {
			return err
		}
		logger.Log.Errorw("failed to check todo existence for update", "id", request.ID, "error", err)
		return errs.NewAppError(500, "failed to validate todo", err)
	}

	updatedTodo := model.Todo{
		ID:     request.ID,
		Todo:   request.Todo,
		UserID: currentUser.ID,
	}

	if err := s.todoRepository.Update(ctx, &updatedTodo); err != nil {
		logger.Log.Errorw("failed to update todo", "id", request.ID, "error", err)
		return errs.NewAppError(500, "failed to update todo", err)
	}

	logger.Log.Infow("todo updated successfully", "id", request.ID)
	return nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, currentUser model.User, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	todo, err := s.todoRepository.GetByID(ctx, id)
	if err != nil {
		if err == errs.ErrTodoNotFound {
			return err
		}
		logger.Log.Errorw("failed to get todo", "id", id, "error", err)
		return errs.NewAppError(500, "failed to get todo", err)
	}

	if todo.UserID != currentUser.ID {
		logger.Log.Errorw("forbidden", "id", id, "error", errs.ErrForbidden)
		return errs.ErrForbidden
	}

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
