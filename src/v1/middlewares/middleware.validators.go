package middlewares

import "github.com/gin-gonic/gin"

func NoValidation(context *gin.Context) {

}

func BasicHeader(context *gin.Context) (msg string, err error) {
	return "", nil
}
