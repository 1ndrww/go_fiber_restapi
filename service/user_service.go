package service

import (
	"example-go-restapi/model/api"
	"example-go-restapi/model/domain"
)

type UserServiceInterface interface {
	FindAll() []domain.User
	FindById(userId string) domain.User
	Save(user api.UserCreateOrUpdateRequest) domain.User
	Edit(user api.UserCreateOrUpdateRequest, userId string) domain.User
	Delete(userId string)
}
