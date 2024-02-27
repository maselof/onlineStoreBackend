package routes

import (
	"net/http"
	"onlineStoreBackend/constants"
	ginhandlers "onlineStoreBackend/entity"
)

func (s RestRoutes) GetAPIHandlers() []ginhandlers.Routes {
	result := []ginhandlers.Routes{
		// users
		{
			Method:      http.MethodPost,
			Route:       constants.RouteUsers,
			HandlerFunc: s.UserController.PostUser,
		},
		{
			Method:      http.MethodDelete,
			Route:       constants.RouteDeleteUser,
			HandlerFunc: s.UserController.DeleteUser,
		},
		{
			Method:      http.MethodGet,
			Route:       constants.RouteUsers,
			HandlerFunc: s.UserController.GetUsers,
		},
		// orders
		{
			Method:      http.MethodPost,
			Route:       constants.RouteOrders,
			HandlerFunc: s.OrderController.PostOrder,
		},
		{
			Method:      http.MethodGet,
			Route:       constants.RouteOrders,
			HandlerFunc: s.OrderController.GetOrdersByUserID,
		},
		{
			Method:      http.MethodDelete,
			Route:       constants.RouteOrders,
			HandlerFunc: s.OrderController.DeleteOrderByUserID,
		},
		// products
		{
			Method:      http.MethodPost,
			Route:       constants.RouteProducts,
			HandlerFunc: s.ProductsController.PostProduct,
		},
		{
			Method:      http.MethodGet,
			Route:       constants.RouteProducts,
			HandlerFunc: s.ProductsController.GetAllProducts,
		},
		{
			Method:      http.MethodDelete,
			Route:       constants.RouteDeleteProduct,
			HandlerFunc: s.ProductsController.DeleteProductByID,
		},
		// cast
		{
			Method:      http.MethodPost,
			Route:       constants.RouteCart,
			HandlerFunc: s.CartsController.PostProductsByUserID,
		},
		{
			Method:      http.MethodGet,
			Route:       constants.RouteCartByUserID,
			HandlerFunc: s.CartsController.GetProductsFromCart,
		},
		{
			Method:      http.MethodDelete,
			Route:       constants.RouteCartByUserID,
			HandlerFunc: s.CartsController.DeleteCast,
		},
	}

	return result
}
