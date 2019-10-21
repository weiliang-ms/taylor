package utils

import (
	"bufio"
	"io"
	"taylor/errors"
	"os"
	"strings"
)

const (
	ContentExist     = "内容已存在"
	AppendSuccessful = "添加成功"
)

// 读取文件关键行内容，返回[]byte类型
func ReadFileKeyWord(filePath string, keyWord string) (keyWordLine string) {
	Logger.Info("读取文件：" + filePath + "关键内容为：" + keyWord)
	f, err := os.Open(filePath)
	defer f.Close()

	if err != nil {
		Logger.Errorln(err.Error())
	}
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if line != "" && strings.Contains(line, keyWord) {
			Logger.Info("命中行数据：" + keyWordLine)
			return line
		}
		if err != nil || err == io.EOF {
			Logger.Errorln(err.Error())
			if keyWordLine == "" {
				keyWordLine = "default"
			}
			break
		}
	}

	return keyWordLine
}

// 读取文件关键行内容，返回[]byte类型
func ModifyContent(filePath string, oldValue string, newValue string) {
	Logger.Info("修改文件：" + filePath + "内容，修改前：" + oldValue + "修改后：" + newValue)
	f, err := os.Open(filePath)
	defer f.Close()

	if err != nil {
		Logger.Errorln(err.Error())
	}
	rd := bufio.NewReader(f)
	out, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0766)
	defer out.Close()

	if err != nil {
		Logger.Errorln("Open write file fail:", err)
		os.Exit(-1)
	}
	for {
		line, err := rd.ReadString('\n')
		if err != nil || err == io.EOF {
			Logger.Errorln(err.Error())
			break
		}
		if re := strings.Contains(line, oldValue); re {
			newLine := strings.Replace(string(line), oldValue, newValue, -1)
			_, err = out.WriteString(newLine + "\n")
			if err != nil {
				Logger.Errorln("write to file fail:", err)
				os.Exit(-1)
			}
		}
	}
}

// 新增文件内容返回结果
func AddLine(filePath string, line string) (err error, msg string) {

	Logger.Info("修改文件：" + filePath + "新增内容：" + line)
	re := ReadFileKeyWord(filePath, line)
	if re == "default" {
		fd, _ := os.OpenFile(filePath, os.O_WRONLY, 0644)
		defer fd.Close()
		seek, _ := fd.Seek(0, 2)
		_, err := fd.WriteAt([]byte("\n"+line), seek)
		if err != nil {
			Logger.Errorln(err.Error())
			return err, AppendSuccessful
		} else {
			return err, AppendSuccessful
		}
	} else {

		return errors.New(ContentExist), ContentExist
	}
}
