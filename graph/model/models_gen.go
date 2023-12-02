// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateTodoInput struct {
	Title string `json:"title"`
}

type Todo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type UpdateTodoInput struct {
	Completed bool `json:"completed"`
}
