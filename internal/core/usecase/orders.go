package usecase

import (
	"onlineStoreBackend/entity/response"
	"onlineStoreBackend/internal/core/repository"
	"onlineStoreBackend/internal/methods"
)

type OrdersService struct {
	repoOrders      repository.OrdersRepository
	ProductsService ProductsService
}

func (s OrdersService) PostOrder(userID int, productsID []int) (result []response.OrdersResponse, err error) {
	tx, err := methods.GetDatabase()
	if err != nil {
		return result, err
	}

	order, err := s.repoOrders.CreateOrder(tx, userID)
	if err != nil {
		return result, err
	}

	ordersProducts := s.AmountProducts(productsID, order.ID)

	result, err = s.repoOrders.PostOrderItems(tx, ordersProducts)
	if err != nil {
		return result, err
	}

	return result, err
}

func (s OrdersService) GetOrdersByUserID(userID int) (result []response.OrdersResponse, err error) {
	tx, err := methods.GetDatabase()
	if err != nil {
		return result, err
	}

	return s.repoOrders.GetOrdersByUserID(tx, userID)
}

func (s OrdersService) AmountProducts(productsIDs []int, orderID int) []response.OrdersResponse {
	helperMap := make(map[int]int)
	for _, val := range productsIDs {
		helperMap[val] += 1
	}

	var orders []response.OrdersResponse
	for key, val := range helperMap {
		orders = append(orders, response.OrdersResponse{
			ID:       orderID,
			Products: key,
			Amount:   val,
		})
	}

	return orders
}

func (s OrdersService) DeleteOrderByUserID(orderID int) (result []response.OrdersResponse, err error) {
	tx, err := methods.GetDatabase()
	if err != nil {
		return result, err
	}

	result, err = s.repoOrders.DeleteOrderItems(tx, orderID)

	err = s.repoOrders.DeleteOrder(tx, orderID)
	if err != nil {
		return result, err
	}

	return result, err
}
