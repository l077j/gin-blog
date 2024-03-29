package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

// getLogFilePath get the log gile save path
func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	// os.Stat：返回文件信息结构描述文件。如果出现错误，会返回*PathError
	_, err := os.Stat(filePath)
	switch {
	// os.IsNotExist：能够接受ErrNotExist、syscall的一些错误，它会返回一个布尔值，能够得知文件不存在或目录不存在
	case os.IsNotExist(err):
		mkDir(getLogFilePath())
	// os.IsPermission：能够接受ErrPermission、syscall的一些错误，它会返回一个布尔值，能够得知权限是否满足
	case os.IsPermission(err):
		log.Fatalf("Permission:%v", err)
	}

	// os.OpenFile：调用文件，支持传入文件名称、指定的模式调用文件、文件权限，返回的文件的方法可以用于I/O。如果出现错误，则为*PathError
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile:%v", err)
	}

	return handle
}

func mkDir(s string) {
	// os.Getwd：返回与当前目录对应的根路径名
	dir, _ := os.Getwd()
	// os.MkdirAll：创建对应的目录以及所需的子目录，若成功则返回nil，否则返回error
	// os.ModePerm：const定义ModePerm FileMode = 0777
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
