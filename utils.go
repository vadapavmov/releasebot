package main

import (
	"io/ioutil"
	"log"
	"net/http"
)


func SendRequest(url string, token string) string {
	
    client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)

    request.Header.Set("Authorization", "Bearer " + token)
    request.Header.Set("Accept","application/json")

	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln("Error reading response body:", err)
	}

	return string(body)
}
