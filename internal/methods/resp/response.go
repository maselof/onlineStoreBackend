package resp

import (
	"github.com/gin-gonic/gin"
	"onlineStoreBackend/entity/response"
)

func SendResponse(c *gin.Context, response response.JSONResponse) {
	c.JSON(response.Code, response)
}
