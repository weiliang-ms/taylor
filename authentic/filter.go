package authentic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 鉴权中间件
func Filter() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("wl_session")
		if err != nil && strings.Contains(c.Request.RequestURI, "/v1/") {
			c.Redirect(http.StatusMovedPermanently, "/static/page/login.html")
		}
	}
}
