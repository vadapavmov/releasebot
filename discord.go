package main

import (
	"encoding/json"
	"fmt"
)


var (
	base string = "https://api.themoviedb.org/3"
)

func SendMessages() string {

	//var u string
	//tv/21212?language=en-US' \
	urlurl := base + "/movie/1219926?language=en-US"
	data := SendRequest(urlurl, TMDB)

	var movie Data

	err := json.Unmarshal([]byte(data), &movie)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	fmt.Println(movie.OriginalTitle)
	fmt.Println(movie.OriginalLanguage)

	return movie.OriginalTitle
}
