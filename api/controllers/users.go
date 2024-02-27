package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"onlineStoreBackend/constants"
	"onlineStoreBackend/entity/request"
	"onlineStoreBackend/entity/response"
	"onlineStoreBackend/internal/core/usecase"
	"onlineStoreBackend/internal/methods/resp"
	"strconv"
)

type UserControllers struct {
	usersService usecase.UsersService
}

func (s UserControllers) PostUser(c *gin.Context) {
	var requestUser request.UserRequest
	buffer, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusBadRequest,
			Message: constants.Errors[constants.ErrorGetParams],
		})

		return
	}

	if err := json.Unmarshal(buffer, &requestUser); err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.Errors[constants.ErrorUnmarshalRequest],
		})

		return
	}

	result, err := s.usersService.PostUser(requestUser)
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

func (s UserControllers) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		resp.SendResponse(c, response.JSONResponse{
			Code:    http.StatusBadRequest,
			Message: constants.Errors[constants.ErrorGetParams],
		})

		return
	}

	result, err := s.usersService.DeleteUser(userIDInt)
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

func (s UserControllers) GetUsers(c *gin.Context) {
	result, err := s.usersService.GetUsers()
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
