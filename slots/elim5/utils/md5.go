package utils

import (
	"crypto/md5"
	"elim5/utils/conver"
	"encoding/hex"
	"github.com/fatih/structs"
	"io"
	"os"
	"sort"
	"strings"
)

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

func Md5Sign(params map[string]interface{}, secret string) string {
	return MD5V([]byte(MakeOriginSign(params, secret)))
}

func MakeOriginSign(params map[string]interface{}, secret string) string {
	var (
		sortedKeys []string
		s          string
		arr        []string
	)
	for k, v := range params {
		params[k] = conver.StringMust(v)
		if k == "sign" || params[k].(string) == "" {
			continue
		}
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	for _, k := range sortedKeys {
		arr = append(arr, k+"="+params[k].(string))
	}
	s = strings.Join(arr, "&") + secret
	return s
}

func StructMd5Sign(params any, secret string) string {
	return Md5Sign(structs.Map(params), secret)
}

func ComputeFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hashInBytes := hash.Sum(nil)[:16]
	return hex.EncodeToString(hashInBytes), nil
}
