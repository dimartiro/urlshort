package services

import (
	"fmt"
	"regexp"

	"github.com/dimartiro/urlshort/api/src/dao"
)

const URL_SIZE = 7

func Short(url string) (string, error) {

	regex := regexp.MustCompile("^(http|https)(://)")

	if !regex.Match([]byte(url)) {
		url = "http://" + url
	}

	hashedUrl := HashUrl(url, URL_SIZE)

	if err := dao.SaveHashUrlInDB(hashedUrl, url); err != nil {
		fmt.Println("Error saving hashed url " + err.Error())
		return "", err
	}

	return hashedUrl, nil
}
