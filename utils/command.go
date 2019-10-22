package utils

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
)

// 执行shell语句
func Shell(shell string) (message string, code int) {

	var out bytes.Buffer
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", shell+" 2>&1")
	default:
		cmd = exec.Command("/bin/bash", "-c", shell+" 2>&1")
	}
	log.Info("执行命令为" + shell)

	cmd.Stdout = &out
	if err := cmd.Start(); err != nil {
		Logger.Errorln(err.Error())
		return err.Error(), 1
	}

	if err := cmd.Wait(); err != nil {
		Logger.Errorln(out.String() + err.Error())
		return err.Error(), 1
	}

	return out.String(), 1
}

// 执行shell语句
func Cmd(shell string) (message string, re error) {

	var out bytes.Buffer
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", shell+" 2>&1")
	default:
		cmd = exec.Command("/bin/bash", "-c", shell+" 2>&1")
	}
	Logger.Info("执行命令为：" + shell)

	cmd.Stdout = &out
	if err := cmd.Start(); err != nil {
		Logger.Errorln("shell start error: ", err.Error())
		re = err
	}

	if err := cmd.Wait(); err != nil {
		Logger.Errorln("shell start error: ",err.Error())
		re = err
	}
	Logger.Info("运行结果：",out.String())
	return out.String(), re
}
