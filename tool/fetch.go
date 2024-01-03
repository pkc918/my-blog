package tool

import (
	"io"
	"log"
	"net/http"
)

func Fetch(url string) []byte {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Http get error is ", err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal("Http status code is ", res.StatusCode)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
