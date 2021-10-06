package service

import (
	"errors"
	"homework-server/contract"
	"homework-server/external"
)

type UserService interface {
	GetDetail(id string) (contract.User, error)
	GetFollower(username string) (int, error)
}

type userServiceImpl struct {
	jsonData map[string]contract.User
}

func NewUserService() UserService {
	return &userServiceImpl{jsonData: external.GetDataJson()}
}

func (u userServiceImpl) GetDetail(id string) (contract.User, error) {
	if user, ok := u.jsonData[id]; ok {
		return user, nil
	} else {
		return contract.User{}, errors.New("not found user")
	}
}

func (u userServiceImpl) GetFollower(username string) (int, error) {
	for _, user := range u.jsonData {
		if user.Username == username {
			return user.Followers, nil
		}
	}
	return 0, errors.New("not found user")
}
