package authentic

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"taylor/utils"
	"time"
)

const (
	LoginSuccessful      = "登录成功"
	GetUserInfoFailed    = "获取用户信息失败"
	BadAccountOrPassword = "账号或密码异常"
)

type LoginInfo struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	var userInfo LoginInfo
	msg := LoginSuccessful
	httpCode := http.StatusOK
	c.BindJSON(&userInfo)

	client := utils.RedisClient()
	defer client.Close()
	re, err := client.Get("user_" + userInfo.UserName).Result()

	if err != nil {
		utils.Logger.Errorln(err.Error())
		httpCode = http.StatusUnauthorized
		msg = GetUserInfoFailed
	}

	if re == "" || re != userInfo.Password {
		httpCode = http.StatusUnauthorized
		msg = BadAccountOrPassword
	}

	if httpCode == http.StatusOK {
		sessionID := generateSessionId()
		result, err := client.Set(sessionID, time.Now().String(), time.Minute*30).Result()
		if err != nil {
			utils.Logger.Errorln(err.Error())
		} else {
			utils.Logger.Info(result)
		}
		//c.SetCookie("session8", GenerateSID(), 3600, "/", "/", true, true)
		c.SetCookie("wl_session", sessionID, 1800, "/", c.Request.Host, false, true)
	}

	c.JSON(httpCode, gin.H{
		"code":     0,
		"error":    nil,
		"msg":      msg,
		"redirect": "/static/index.html",
	})

}

// GenerateSID 产生唯一的Session ID
func GenerateSID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
