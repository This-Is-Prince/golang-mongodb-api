package services

import "github.com/This-Is-Prince/golang-mongodb-api/models"

type UserService interface {
	FindUserById(string) (*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse, error)
}
