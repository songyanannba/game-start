package helper

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func CutPathByFirstPart(path string) (string, string) {
	path = strings.TrimLeft(path, "/")
	index := strings.Index(path, "/")
	if index != -1 {
		return path[:index], path[index+1:]
	}
	return path, ""
}

func MoveFile(src, dst string) (err error) {
	// 确保目的地目录不存在
	if _, err = os.Stat(dst); !os.IsNotExist(err) {
		return errors.New("destination directory already exists: " + err.Error())
	}

	// 获取源文件夹的绝对路径
	src, err = filepath.Abs(src)
	if err != nil {
		return errors.New("Failed to get the absolute path of the source folder: " + err.Error())
	}

	// 创建目的地文件夹
	err = os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	if err != nil {
		return errors.New("Failed to create destination folder: " + err.Error())
	}

	// 移动文件夹
	err = os.Rename(src, dst)
	if err != nil {
		return errors.New("Failed to move folder: " + err.Error())
	}
	return nil
}
