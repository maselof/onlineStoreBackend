package entity

import "github.com/gin-gonic/gin"

type Routes struct {
	Method      string
	Route       string
	HandlerFunc gin.HandlerFunc
}
