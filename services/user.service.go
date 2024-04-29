package services

import "example.com/go-practice/models"


type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	DeleteUser(*string) error
	UpdateUser(*models.User) error
}