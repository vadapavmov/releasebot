package main

import (
	"encoding/json"
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

func MapToJSONString(inputMap map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(inputMap)

	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}


/*
func Parser(body string) string {

	var response Response

	err := json.Unmarshal([]byte(body), &response)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}

	if len(response.Choices) > 0 {
		content := response.Choices[0].Message.Content
		return content
	}

	return ""
}
*/
