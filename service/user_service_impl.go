package service

import (
	"example-go-restapi/exception"
	"example-go-restapi/model/api"
	"example-go-restapi/model/domain"
	"example-go-restapi/repository"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserService(userRepository repository.UserRepositoryInterface) UserServiceInterface {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (s UserServiceImpl) FindAll() []domain.User {
	return s.UserRepository.FindAll()
}

func (s UserServiceImpl) FindById(userId string) domain.User {
	user, err := s.UserRepository.FindById(userId)
	if err != nil {
		log.Warn(err)
		panic(exception.NewNotFoundError(err.Error()))
	}
	return user
}

func (s UserServiceImpl) Save(user api.UserCreateOrUpdateRequest) domain.User {
	userDomain := domain.User{
		ID:        uuid.NewString(),
		UserName:  user.UserName,
		Email:     user.Email,
		Phone:     user.Phone,
		FullName:  user.FullName,
		CreatedAt: strconv.Itoa(int(time.Now().Unix())),
	}
	return s.UserRepository.Save(userDomain)
}

func (s UserServiceImpl) Edit(user api.UserCreateOrUpdateRequest, userId string) domain.User {
	u, err := s.UserRepository.FindById(userId)
	if err != nil {
		log.Warn(err)
		panic(exception.NewNotFoundError(err.Error()))
	}
	userDomain := domain.User{
		ID:        userId,
		UserName:  user.UserName,
		Email:     user.Email,
		Phone:     user.Phone,
		FullName:  user.FullName,
		UpdatedAt: strconv.Itoa(int(time.Now().Unix())),
		CreatedAt: u.CreatedAt,
	}

	u, err = s.UserRepository.Edit(userDomain)
	if err != nil {
		log.Warn(err)
		panic(exception.NewNotFoundError(err.Error()))
	}
	return u
}

func (s UserServiceImpl) Delete(userId string) {
	err := s.UserRepository.Delete(userId)
	if err != nil {
		log.Warn(err)
		panic(exception.NewNotFoundError(err.Error()))
	}
}
