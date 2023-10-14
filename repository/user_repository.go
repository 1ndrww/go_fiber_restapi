package repository

import (
	"example-go-restapi/model/domain"
)

type UserRepositoryInterface interface {
	Save(user domain.User) domain.User
	FindById(userId string) (domain.User, error)
	FindAll() []domain.User
	Edit(user domain.User) (domain.User, error)
	Delete(userId string) error
}
