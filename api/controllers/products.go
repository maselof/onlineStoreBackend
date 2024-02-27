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

type ProductsController struct {
	productsService usecase.ProductsService
}

func (s ProductsController) PostProduct(c *gin.Context) {
	var requestProduct request.ProductRequest

	buffer, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusBadRequest,
			Message: constants.Errors[constants.ErrorGetParams],
		})

		return
	}

	if err := json.Unmarshal(buffer, &requestProduct); err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.Errors[constants.ErrorUnmarshalRequest],
		})

		return
	}

	result, err := s.productsService.PostProduct(requestProduct)
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

func (s ProductsController) GetAllProducts(c *gin.Context) {
	result, err := s.productsService.GetAllProducts()
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

func (s ProductsController) DeleteProductByID(c *gin.Context) {
	strID := c.Param("id")

	intID, err := strconv.Atoi(strID)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusBadRequest,
			Message: constants.Errors[constants.ErrorGetParams],
		})

		return
	}

	result, err := s.productsService.DeleteProductByID(intID)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal DB",
			Data:    err,
		})

		log.Fatal(err)

		return
	}

	resp.SendResponse(c, response.JSONResponse{
		Code:    http.StatusOK,
		Message: "Data deleted",
		Data:    result,
	})
}
