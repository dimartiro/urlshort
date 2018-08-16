package services

import (
	"regexp"

	"github.com/dimartiro/urlshort/model"
)

const URL_SIZE = 7

func Short(url string) string {

	regex, _ := regexp.Compile("^(http|https)(://)")

	if !regex.Match([]byte(url)) {
		url = "http://" + url
	}

	hashedUrl := HashUrl(url, URL_SIZE)

	model.SaveHashToUrl(hashedUrl, url)

	return hashedUrl
}
