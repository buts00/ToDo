package service

import (
	Todo "github.com/buts00/ToDo"
	"github.com/buts00/ToDo/pkg/repository"
)

type Authorization interface {
	CreateNewUser(user Todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ToDoList interface {
}

type ToDoItem interface {
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
