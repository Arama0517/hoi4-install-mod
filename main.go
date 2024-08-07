package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/go-cmd/cmd"
)

const version = "devel"

const (
	ErrInitFail = iota + 1
)

func init() {
	// 检测是否是 Windows 系统
	if runtime.GOOS != "windows" {
		log.Fatal("此程序仅支持Windows系统")
		os.Exit(ErrInitFail)
	}

	// 设置 log 的样式
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)

	// DLL 文件路径
	currentPath, err := os.Getwd()
	if err != nil {
		log.WithError(err).Fatal("获取当前目录失败")
		os.Exit(ErrInitFail)
	}
	if fileExist(filepath.Join(currentPath, "steam_api.dll")) || fileExist(filepath.Join(currentPath, "steam_api64.dll")) {
		if fileExist(filepath.Join(currentPath, "main.ps1")) {
			c := cmd.NewCmd("powershell", "-Command", "Invoke-WebRequest -Uri https://mirror.ghproxy.com/https://raw.githubusercontent.com/Arama0517/hoi4-install-mod/main/main.ps1 -OutFile main.ps1")
			c.Dir = currentPath
			statusChan := c.Start()
			status := <-statusChan
			if status.Exit != 0 {
				log.WithError(status.Error).Fatal("下载脚本失败")
			}
		}
		c := cmd.NewCmd("powershell", "-File", filepath.Join(currentPath, "main.ps1"))
		c.Dir = currentPath
		statusChan := c.Start()
		status := <-statusChan
		if status.Exit != 0 {
			log.WithError(status.Error).Fatal("运行脚本失败")
		}
	}
}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func main() {
	// todo: 下载mod..
}
