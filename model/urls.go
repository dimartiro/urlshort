package model

var urls map[string]string = make(map[string]string)

func SaveHashToUrl(hash, url string) {
	urls[hash] = url
}

func GetUrlFromHash(hash string) string {
	return urls[hash]
}
