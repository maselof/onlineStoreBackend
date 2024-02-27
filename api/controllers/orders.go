package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineStoreBackend/constants"
	"onlineStoreBackend/entity/response"
	"onlineStoreBackend/internal/core/usecase"
	"onlineStoreBackend/internal/methods/resp"
	"strconv"
)

type OrderController struct {
	orderService usecase.OrdersService
	castService  usecase.CartsService
}

func (s OrderController) PostOrder(c *gin.Context) {
	strID := c.Param("id")

	intID, err := strconv.Atoi(strID)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusBadRequest,
			Message: constants.Errors[constants.ErrorGetParams],
		})

		return
	}

	productsIDs, err := s.castService.GetProductsFromCart(c, intID)
	if len(productsIDs) == 0 {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusNoContent,
			Message: constants.Errors[constants.ErrorEmptyCast],
		})

		return
	}

	result, err := s.orderService.PostOrder(intID, productsIDs)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal DB",
			Data:    err,
		})

		return
	}

	resp.SendResponse(c, response.JSONResponse{
		Code: http.StatusOK,
		Data: result,
	})
}

func (s OrderController) GetOrdersByUserID(c *gin.Context) {
	userID := c.Param("id")

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusBadRequest,
			Message: constants.Errors[constants.ErrorGetParams],
		})

		return
	}

	result, err := s.orderService.GetOrdersByUserID(intUserID)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal DB",
			Data:    err,
		})

		return
	}

	resp.SendResponse(c, response.JSONResponse{
		Code: http.StatusOK,
		Data: result,
	})
}

func (s OrderController) DeleteOrderByUserID(c *gin.Context) {
	orderID := c.Param("id")

	intOrderID, err := strconv.Atoi(orderID)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusBadRequest,
			Message: constants.Errors[constants.ErrorGetParams],
		})

		return
	}

	result, err := s.orderService.DeleteOrderByUserID(intOrderID)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal DB",
			Data:    err,
		})

		return
	}

	resp.SendResponse(c, response.JSONResponse{
		Code: http.StatusOK,
		Data: result,
	})
}
