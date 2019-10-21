package nginx

import (
	"io/ioutil"
	"taylor/utils"
	"os"
)

// 读取文件内容，返回[]byte类型
func readFileContent(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		utils.Logger.Errorln(err)
		return nil
	}
	return b
}

// 写入自定义配置信息
func writeConfigContent(path string, content string) (msg string, code int) {
	// 写入配置文件内容
	file, err := os.Create(path)
	defer file.Close()

	if err != nil {
		msg = err.Error()
		os.Remove(path)
	} else {
		_, writeErr := file.Write([]byte(content))
		if writeErr != nil {
			msg = writeErr.Error()
		} else {
			// 检测配置内容是否合法
			msg, code = check()
		}
	}
	return msg, 0
}
