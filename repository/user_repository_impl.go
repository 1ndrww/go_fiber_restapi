package repository

import (
	"errors"
	"example-go-restapi/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepositoryImpl{}
}

var Users []domain.User

func (u UserRepositoryImpl) Save(user domain.User) domain.User {
	Users = append(Users, user)
	return user
}

func (u UserRepositoryImpl) FindById(userId string) (domain.User, error) {
	for _, user := range Users {
		if user.ID == userId {
			return user, nil
		}
	}
	return domain.User{}, errors.New("user id is not found")
}

func (u UserRepositoryImpl) FindAll() []domain.User {
	return Users
}

func (u UserRepositoryImpl) Edit(user domain.User) (domain.User, error) {
	for _, u := range Users {
		if u.ID == user.ID {
			u.Email = user.Email
			u.FullName = user.FullName
			u.Phone = user.Phone
			u.UserName = user.UserName
			return user, nil
		}
	}

	return domain.User{}, errors.New("user id is not found")
}

func (u UserRepositoryImpl) Delete(userId string) error {
	var isFindById bool
	var indexUser int

	for i, u := range Users {
		if u.ID == userId {
			isFindById = true
			indexUser = i
		}
	}
	if isFindById == false {
		return errors.New("user id is not found")
	}
	Users = append(Users[:indexUser], Users[indexUser+1:]...)

	return nil
}
