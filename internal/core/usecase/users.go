package usecase

import (
	"onlineStoreBackend/entity/request"
	"onlineStoreBackend/entity/response"
	"onlineStoreBackend/internal/core/repository"
	"onlineStoreBackend/internal/methods"
)

type UsersService struct {
	repoUsers repository.UsersRepository
}

func (s UsersService) PostUser(requestUser request.UserRequest) (result response.UsersResponse, err error) {
	tx, err := methods.GetDatabase()
	if err != nil {
		return result, err
	}
	defer tx.Close()

	return s.repoUsers.PostUser(tx, requestUser.Name)
}

func (s UsersService) DeleteUser(id int) (result response.UsersResponse, err error) {
	tx, err := methods.GetDatabase()
	if err != nil {
		return result, err
	}
	defer tx.Close()

	return s.repoUsers.DeleteUser(tx, id)
}

func (s UsersService) GetUsers() (result []response.UsersResponse, err error) {
	tx, err := methods.GetDatabase()
	if err != nil {
		return result, err
	}
	defer tx.Close()

	return s.repoUsers.GetUsers(tx)
}
