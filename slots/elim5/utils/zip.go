package utils

import (
	"archive/zip"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

type GameFile struct {
	FullName string
	MD5      string
}

// CreateZipFileAndMd5 创建文件并将文件名称加入md5 返回文件路径
func CreateZipFileAndMd5(f *zip.File, destDir, fileName string) (string, error) {
	// 创建文件夹
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return "", err
	}

	inFile, err := f.Open()
	if err != nil {
		return "", err
	}
	defer inFile.Close()

	hash := md5.New()

	filePath := filepath.Join(destDir, fileName)
	outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// 使用 TeeReader 在复制数据到 outFile 时同时写入到哈希
	tee := io.TeeReader(inFile, hash)

	_, err = io.Copy(outFile, tee)
	if err != nil {
		return "", err
	}

	// 在数据被读取后计算 MD5 值
	hashInBytes := hash.Sum(nil)[:16]
	md5Str := hex.EncodeToString(hashInBytes)

	// 可以根据需要将 MD5 值加到文件名中
	newFilePath := filepath.Join(destDir, md5Str+fileName)
	os.Rename(filePath, newFilePath)

	return filePath, nil
}

func CreateZipFile(f *zip.File, destDir, fileName string) (string, error) {
	// 创建文件夹
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return "", err
	}

	inFile, err := f.Open()
	if err != nil {
		return "", err
	}
	defer inFile.Close()

	// 拼接文件名
	fPath := filepath.Join(destDir, fileName)
	// 写文件
	outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, inFile)
	if err != nil {
		return "", err
	}

	return fPath, nil
}
