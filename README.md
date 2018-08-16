# URLSHORT
URL shortener made in go

## Technologies
* Go 1.10.3
* Nginx 1.15.2

## Setup
The project is fully setup using Docker Compose.

Run it using:
```
docker-compose build
docker-compose up -d
```

## Use

Short URL

```
curl -X POST 'http://localhost/short?url=google.com'
```

Get URL from shorted link
```
curl -X GET 'http://localhost/Yml0Lmx'
```
