package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"onlineStoreBackend/constants"
	"onlineStoreBackend/entity/request"
	"onlineStoreBackend/entity/response"
	"onlineStoreBackend/internal/core/usecase"
	"onlineStoreBackend/internal/methods/resp"
	"strconv"
)

type CartsController struct {
	cartsService usecase.CartsService
}

func (s CartsController) PostProductsByUserID(c *gin.Context) {
	var requestCast request.CastRequest

	buffer, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusBadRequest,
			Message: constants.Errors[constants.ErrorGetParams],
		})

		return
	}

	if err := json.Unmarshal(buffer, &requestCast); err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.Errors[constants.ErrorUnmarshalRequest],
		})

		return
	}

	err = s.cartsService.PostProducts(c, requestCast)
	log.Print(err)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal DB",
			Data:    err.Error(),
		})

		return
	}

	resp.SendResponse(c, response.JSONResponse{
		Code:    http.StatusOK,
		Message: "Data write successfully",
	})
}

func (s CartsController) GetProductsFromCart(c *gin.Context) {
	strID := c.Param("id")

	intID, err := strconv.Atoi(strID)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusBadRequest,
			Message: constants.Errors[constants.ErrorGetParams],
		})

		return
	}

	result, err := s.cartsService.GetProductsFromCart(c, intID)
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

func (s CartsController) DeleteCast(c *gin.Context) {
	strID := c.Param("id")

	strInt, err := strconv.Atoi(strID)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusBadRequest,
			Message: constants.Errors[constants.ErrorGetParams],
		})

		return
	}

	err = s.cartsService.DeleteCart(c, strInt)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal DB",
			Data:    err,
		})

		return
	}

	resp.SendResponse(c, response.JSONResponse{
		Code:    http.StatusOK,
		Message: "Data delete successfully",
	})
}
