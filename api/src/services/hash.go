package services

import (
	"crypto/md5"
	"encoding/base64"
	"regexp"
)

func HashUrl(url string, size int) string {
	regex := regexp.MustCompile("[.|-|+|=]")

	hashSum := md5.Sum([]byte(url))
	base64Encoded := base64.StdEncoding.EncodeToString(hashSum[:])
	hash := string(regex.ReplaceAll([]byte(base64Encoded), []byte("")))

	return hash[:size]
}
