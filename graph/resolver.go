package graph

import (
	"context"

	"github.com/bryanpinheiro/go-graphql-todoapi/database"
	"github.com/bryanpinheiro/go-graphql-todoapi/models"
)

type Resolver struct{}

type CreateTodoInput struct {
	Title string
}

type UpdateTodoInput struct {
	Completed bool
}

// GET ALL TODOS
func (r *Resolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	// Access the DB instance from the database package
	db := database.DB

	var todos []*models.Todo
	if err := db.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

// GET TODO
func (r *Resolver) Todo(ctx context.Context, args struct{ ID uint }) (*models.Todo, error) {
	var todo models.Todo

	// Access the DB instance from the database package
	db := database.DB

	// Fetch the todo by ID from the database
	if err := db.First(&todo, args.ID).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

// CreateTodo resolver for creating a new todo
func (r *Resolver) CreateTodo(ctx context.Context, input CreateTodoInput) (*models.Todo, error) {
	todo := models.Todo{
		Title: input.Title,
	}
	database.DB.Create(&todo)
	return &todo, nil
}

// UpdateTodo resolver for updating an existing todo
func (r *Resolver) UpdateTodo(ctx context.Context, id uint, input UpdateTodoInput) (*models.Todo, error) {
	var todo models.Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}

	todo.Completed = input.Completed

	database.DB.Save(&todo)
	return &todo, nil
}

// DeleteTodo resolver for deleting a todo
func (r *Resolver) DeleteTodo(ctx context.Context, id uint) (bool, error) {
	var todo models.Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		return false, err
	}

	database.DB.Delete(&todo)
	return true, nil
}
