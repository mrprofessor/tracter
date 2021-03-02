package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Read the response body and return as string
func ParseResponse(resp *http.Response) string {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}
