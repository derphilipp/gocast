package cmd

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func loadBodyfromURL(url string) []byte {
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "gocast")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}
