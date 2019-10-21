package utils

import (
	"os"
	"runtime"
	"strings"
)

// 获取目录路径，并删除结尾的"/"（windows下为"\"）
func CurrentDir(path string) string {
	fileInfo, err := os.Stat(path)
	if err != nil {
		Logger.Errorln(err)
	}
	if fileInfo.IsDir() {
		switch runtime.GOOS {
		case "windows":
			return strings.TrimSuffix(path, "\\")
		default:
			return strings.TrimSuffix(path, "/")
		}
	}
	switch runtime.GOOS {
	case "windows":
		return strings.TrimSuffix(path, "\\"+fileInfo.Name())
	default:
		return strings.TrimSuffix(path, "/"+fileInfo.Name())
	}
}

// '/'的格式
func slashFormat() string {
	switch runtime.GOOS {
	case "windows":
		return "\\"
	default:
		return "/"
	}
}

// 获取目录路径，结尾拼接"/"（windows下为"\"）
func DirAppendSlash(path string) (re string) {

	dir, _ := os.Stat(path)
	endWithSlash := strings.HasSuffix(path, "/") || strings.HasSuffix(path, "\\")
	if dir.IsDir() && endWithSlash {
		return path
	}
	return path + slashFormat()
}

// 拼接路径,删除最后一个'/'或者'\'
func AppendPathWithSlash(path ...string) (re string) {
	for _, v := range path {
		re = DirAppendSlash(re + v)
	}
	return strings.TrimSuffix(re, slashFormat())
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
