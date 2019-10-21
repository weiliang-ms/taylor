package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/sirupsen/logrus"
)

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为
	// fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	re := hex.EncodeToString(h.Sum(nil))
	logrus.Info(str, " 加密结果为：", re)
	return re
}
