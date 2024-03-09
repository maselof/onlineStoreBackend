package usecase

import (
	"onlineStoreBackend/entity/request"
	"onlineStoreBackend/entity/response"
	"onlineStoreBackend/internal/core/repository"
	"onlineStoreBackend/internal/methods"
	"strconv"
	"strings"
)

type ProductsService struct {
	repoProducts repository.ProductsRepository
	repoOrders   repository.OrdersRepository
}

func (s ProductsService) PostProduct(request request.ProductRequest) (result response.ProductsResponse, err error) {
	tx, err := methods.GetDatabase()
	if err != nil {
		return result, err
	}

	data, err := s.repoProducts.PostProduct(tx, request)

	result.ID = data.ID
	result.Name = data.Name

	data.Price = strings.Replace(data.Price, "₽", "", -1)
	result.Price, err = strconv.ParseFloat(data.Price, 64)
	return
}

func (s ProductsService) GetAllProducts() (result []response.ProductsResponse, err error) {
	tx, err := methods.GetDatabase()
	if err != nil {
		return result, err
	}

	data, err := s.repoProducts.GetAllProducts(tx)
	if err != nil {
		return result, err
	}

	result = make([]response.ProductsResponse, len(data))

	for ind, val := range data {
		result[ind].ID = val.ID
		result[ind].Name = val.Name

		val.Price = strings.Replace(val.Price, "₽", "", -1)

		result[ind].Price, err = strconv.ParseFloat(val.Price, 64)
		if err != nil {
			return result, err
		}
	}

	return result, err
}

func (s ProductsService) GetProductsByIds(ids []int) (result []response.ProductsResponse, err error) {
	tx, err := methods.GetDatabase()
	if err != nil {
		return result, err
	}

	data, err := s.repoProducts.GetProductsByIds(tx, ids)
	if err != nil {
		return result, err
	}

	result = make([]response.ProductsResponse, len(data))

	for ind, val := range data {
		result[ind].ID = val.ID
		result[ind].Name = val.Name

		val.Price = strings.Replace(val.Price, "₽", "", -1)

		result[ind].Price, err = strconv.ParseFloat(val.Price, 64)
		if err != nil {
			return result, err
		}
	}

	return result, err
}

func (s ProductsService) DeleteProductByID(id int) (result response.ProductsResponse, err error) {
	tx, err := methods.GetDatabase()
	if err != nil {
		return result, err
	}

	err = s.repoOrders.DeleteOrderItemByItemID(tx, id)
	if err != nil {
		return result, err
	}

	data, err := s.repoProducts.DeleteProductByID(tx, id)
	if err != nil {
		return result, err
	}

	result.ID = data.ID
	result.Name = data.Name

	data.Price = strings.Replace(data.Price, "₽", "", -1)
	result.Price, err = strconv.ParseFloat(data.Price, 64)
	if err != nil {
		return result, err
	}

	return result, err
}
