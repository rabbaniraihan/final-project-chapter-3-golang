package service

import (
	"final-project/helper"
	"final-project/model"
	"final-project/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) Register(userRegisterRequest model.UserRegisterRequest) (string, error) {
	id := helper.GenerateID()
	hashPassword, err := helper.Hash(userRegisterRequest.Password)
	if err != nil {
		return "", err
	}

	user := model.User{
		Id:       id,
		Username: userRegisterRequest.Username,
		Password: hashPassword,
	}

	err = us.UserRepository.Add(user)
	if err != nil {
		return "", nil
	}

	return id, nil
}

func (us *UserService) Login(userLoginRequest model.UserLoginRequest) (string, error) {
	user, err := us.UserRepository.GetByUsername(userLoginRequest.Username)
	if err != nil {
		return "", err
	}

	if !helper.IsHashValid(user.Password, userLoginRequest.Password) {
		return "", model.ErrorInvalidUsernameOrPassword
	}

	token, err := helper.GenerateToken(user.Id)
	return token, err
}
