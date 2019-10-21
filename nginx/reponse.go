package nginx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseType uint

const (
	Successful    = ResponseType(iota) // 200
	Unauthorized                       // 401
	Forbidden                          // 403
	NotFound                           // 404
	InternalError                      // 500
	Unknown
)

func response(c *gin.Context, responseType ResponseType, msg string) {
	switch responseType {
	case Successful:
		successfulResponse(c, msg)
	case Forbidden:
		forbiddenResponse(c, msg)
	case InternalError:
		internalServerErrorResponse(c, msg)
	default:
		successfulResponse(c, msg)
	}
}

// 200响应
func successfulResponse(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
	})
}

// 403响应
func forbiddenResponse(c *gin.Context, msg string) {
	c.JSON(http.StatusForbidden, gin.H{
		"code": 0,
		"msg":  msg,
	})
}

// 500响应
func internalServerErrorResponse(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 0,
		"msg":  msg,
	})
}
