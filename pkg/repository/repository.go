package repository

import (
	Todo "github.com/buts00/ToDo"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateNewUser(user Todo.User) (int, error)
	User(username, password string) (Todo.User, error)
}

type ToDoList interface {
}

type ToDoItem interface {
}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
