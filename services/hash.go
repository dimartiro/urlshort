package services

import (
	"crypto/md5"
	"encoding/base64"
	"regexp"
)

func HashUrl(url string, size int) string {
	h := md5.New()
	hashSum := h.Sum([]byte(url))
	regex, _ := regexp.Compile("[.|-|+|=]")

	hash := string(regex.ReplaceAll([]byte(base64.StdEncoding.EncodeToString(hashSum)), []byte("")))

	return hash[:size]
}
