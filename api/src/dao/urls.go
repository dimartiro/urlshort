package dao

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var pgConnStr string = os.Getenv("DB_CONN_STR")
var pgHost string = os.Getenv("DB_HOST")
var pgUser string = os.Getenv("DB_USER")
var pgDb string = os.Getenv("DB_DATABASE")
var dbinfo string = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", pgHost, pgUser, pgDb)

func SaveHashUrlInDB(hash, url string) error {
	db, err := sql.Open("postgres", pgConnStr)

	if err != nil {
		return err
	}

	insert := fmt.Sprintf("INSERT INTO links (hash, url) VALUES('%s', '%s')", hash, url)
	_, err = db.Exec(insert)

	if err != nil {
		return err
	}

	err = db.Close()

	if err != nil {
		return err
	}

	return nil
}

func GetUrlFromHashFromDB(hash string) (string, error) {
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return "", err
	}

	var url string
	err = db.QueryRow("SELECT url FROM links WHERE hash = $1", hash).Scan(&url)

	if err != nil {
		return "", err
	}

	return url, nil
}
