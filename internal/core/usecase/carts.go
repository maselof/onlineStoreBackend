package usecase

import (
	"context"
	"fmt"
	"onlineStoreBackend/entity/request"
	"onlineStoreBackend/internal/core/repository"
	"onlineStoreBackend/internal/methods"
	"strconv"
)

type CartsService struct {
	repoCarts    repository.CartsRepository
	repoUsers    repository.UsersRepository
	repoProducts repository.ProductsRepository
}

func (s CartsService) PostProducts(ctx context.Context, request request.CastRequest) (err error) {
	client, err := methods.GetRedis(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	tx, err := methods.GetDatabase()
	if err != nil {
		return err
	}

	user, err := s.repoUsers.GetUserByID(tx, request.UserID)
	if err != nil {
		return err
	}

	if len(user) == 0 {
		return fmt.Errorf("user with id %d doesn't exist", request.UserID)
	}

	products, err := s.repoProducts.GetProductsByIds(tx, request.ProductsID)
	if err != nil {
		return err
	}

	helperMap := make(map[int]int)
	for _, val := range request.ProductsID {
		helperMap[val] = 1
	}

	if len(helperMap) != len(products) {
		return fmt.Errorf("products doesn't exist")
	}

	err = s.repoCarts.DeleteCart(ctx, client, request.UserID)
	if err != nil {
		return err
	}

	productsIds := make([]interface{}, len(request.ProductsID))
	for ind, val := range request.ProductsID {
		productsIds[ind] = val
	}

	return s.repoCarts.PostProducts(ctx, client, request.UserID, productsIds)
}

func (s CartsService) DeleteCart(ctx context.Context, userID int) (err error) {
	client, err := methods.GetRedis(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	return s.repoCarts.DeleteCart(ctx, client, userID)
}

func (s CartsService) GetProductsFromCart(ctx context.Context, userID int) (ids []int, err error) {
	client, err := methods.GetRedis(ctx)
	if err != nil {
		return ids, err
	}
	defer client.Close()

	data, err := s.repoCarts.GetProductsFromCart(ctx, client, userID)
	if err != nil {
		return ids, err
	}

	ids = make([]int, len(data))

	for ind, val := range data {
		ids[ind], err = strconv.Atoi(val)
		if err != nil {
			return ids, err
		}
	}

	return ids, err
}
