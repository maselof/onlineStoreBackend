package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineStoreBackend/api/controllers"
)

type RestRoutes struct {
	OrderController    controllers.OrderController
	UserController     controllers.UserControllers
	ProductsController controllers.ProductsController
	CartsController    controllers.CartsController
}

func (s RestRoutes) Setup(group *gin.RouterGroup) {
	handlerFunc := s.GetAPIHandlers()

	for _, val := range handlerFunc {
		var registerFunc func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

		switch val.Method {
		case http.MethodPost:
			registerFunc = group.POST
		case http.MethodGet:
			registerFunc = group.GET
		case http.MethodDelete:
			registerFunc = group.DELETE
		case http.MethodPut:
			registerFunc = group.PUT
		case http.MethodPatch:
			registerFunc = group.PATCH
		default:
			registerFunc = group.GET
		}

		registerFunc(val.Route, val.HandlerFunc)
	}

}
