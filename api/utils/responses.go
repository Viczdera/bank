package utils

import "github.com/gin-gonic/gin"

func ErrResponse(err error) gin.H {
	return gin.H{"res error": err.Error()}
}
