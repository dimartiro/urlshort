package services

import "github.com/dimartiro/urlshort/model"

const URL_SIZE = 7

func Short(url string) string {
	hashedUrl := HashUrl(url, URL_SIZE)
	model.SaveHashToUrl(hashedUrl, url)

	return hashedUrl
}
