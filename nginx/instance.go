package nginx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taylor/utils"
)

// 检测配置文件是否合法
func check() (msg string, code int) {
	cmd := BinaryPath() + CheckConfigDirect
	return utils.Shell(cmd)
}

// 启动
func Start(c *gin.Context) {
	cmd := BinaryPath()
	msg, code := utils.Shell(cmd)
	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"direct": "start",
		"msg":    msg,
	})
}

// 运行状态
func Status() (msg string, code int) {
	cmd := "ps -ef|grep \"nginx: master\"|grep -v grep"
	return utils.Shell(cmd)
}

// 关闭
func Stop(c *gin.Context) {
	cmd := BinaryPath() + StopInstanceDirect
	msg, code := utils.Shell(cmd)
	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"direct": "stop",
		"msg":    msg,
	})
}

// 平滑关闭
func Quit() (msg string, code int) {
	cmd := BinaryPath() + QuitInstanceDirect
	return utils.Shell(cmd)
}

// 重载配置
func Reload(c *gin.Context) {
	cmd := BinaryPath() + ReloadConfigDirect
	msg, err := utils.Cmd(cmd)
	if err != nil {
		response(c, InternalError, msg)
	} else {
		response(c, Successful, "重载配置成功！")
	}
}
