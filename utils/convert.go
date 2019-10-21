package utils

import (
	"net/http"
	"runtime"
)

func ConvertHttpStatus(err error) int {
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

// TODO:后续调整方法至其他文件内
func Linefeed() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}
