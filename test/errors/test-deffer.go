package main

import (
	"fmt"
	"net/http"
)

func callURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// Make sure we close the Body
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")

	if contentType == "" {
		return "", fmt.Errorf("Can't find the content-type header")
	}

	return contentType, nil
}

func main() {
	cntType, err := callURL("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	} else {
		fmt.Println(cntType)
	}
}
