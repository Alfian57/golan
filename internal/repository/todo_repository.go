package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Alfian57/belajar-golang/internal/database"
	"github.com/Alfian57/belajar-golang/internal/dto"
	errs "github.com/Alfian57/belajar-golang/internal/errors"
	"github.com/Alfian57/belajar-golang/internal/model"
	"github.com/Alfian57/belajar-golang/internal/utils/queryBuilder"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TodoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{db: database.DB}
}

func (r *TodoRepository) getTodos(ctx context.Context, query string, args ...any) ([]model.Todo, error) {
	todos := []model.Todo{}
	err := r.db.SelectContext(ctx, &todos, query, args...)
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (r *TodoRepository) GetAll(ctx context.Context, queryParam dto.GetTodosFilter) ([]model.Todo, error) {
	baseQuery := "SELECT id, todo, user_id, created_at, updated_at FROM todos"

	qb := queryBuilder.NewQueryBuilder(baseQuery)
	qb.Search("todo", queryParam.Search).
		OrderBy(queryParam.OrderBy, queryParam.OrderType).
		Paginate(queryParam.PaginationRequest)

	query, args := qb.Build()
	return r.getTodos(ctx, query, args...)
}

func (r *TodoRepository) GetAllByUser(ctx context.Context, userID string, queryParam dto.GetTodosFilter) ([]model.Todo, error) {
	baseQuery := "SELECT id, todo, user_id, created_at, updated_at FROM todos"

	qb := queryBuilder.NewQueryBuilder(baseQuery)
	qb.Where("user_id = ?", userID).
		Search("todo", queryParam.Search).
		OrderBy(queryParam.OrderBy, queryParam.OrderType).
		Paginate(queryParam.PaginationRequest)

	query, args := qb.Build()
	return r.getTodos(ctx, query, args...)
}

func (r *TodoRepository) CountAll(ctx context.Context, queryParam dto.GetTodosFilter) (int64, error) {
	baseQuery := "SELECT COUNT(*) FROM todos"

	qb := queryBuilder.NewQueryBuilder(baseQuery)
	qb.Search("todo", queryParam.Search)

	query, args := qb.BuildCount(baseQuery)

	var count int64
	err := r.db.GetContext(ctx, &count, query, args...)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *TodoRepository) CountAllByUser(ctx context.Context, userID string, queryParam dto.GetTodosFilter) (int64, error) {
	baseQuery := "SELECT COUNT(*) FROM todos"

	qb := queryBuilder.NewQueryBuilder(baseQuery)
	qb.Where("user_id = ?", userID).
		Search("todo", queryParam.Search)

	query, args := qb.BuildCount(baseQuery)

	var count int64
	err := r.db.GetContext(ctx, &count, query, args...)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *TodoRepository) GetByID(ctx context.Context, id string) (model.Todo, error) {
	todo := model.Todo{}
	query := "SELECT id, todo, user_id, created_at, updated_at FROM todos WHERE id = ?"

	err := r.db.GetContext(ctx, &todo, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return todo, errs.ErrTodoNotFound
		}
		return todo, err
	}

	return todo, nil
}

func (r *TodoRepository) Create(ctx context.Context, todo *model.Todo) error {
	query := "INSERT INTO todos(id, todo, user_id) VALUES (?, ?, ?)"

	_, err := r.db.ExecContext(ctx, query, uuid.New().String(), todo.Todo, todo.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) Update(ctx context.Context, todo *model.Todo) error {
	query := "UPDATE todos SET todo = ?, user_id = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, todo.Todo, todo.UserID, todo.ID)
	return err
}

func (r *TodoRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM todos WHERE id = ?"
	result, err := r.db.ExecContext(ctx, query, id)

	rowsAffected, err := result.RowsAffected()

	if rowsAffected == 0 {
		return errs.ErrTodoNotFound
	}

	return err
}
