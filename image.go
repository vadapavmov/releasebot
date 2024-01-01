package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"
	"net/http"
	"os"
)

func JPEGHandler(res *http.Response) ([]byte, error) {
	// Open the image file
	filePath := "./WARS.jpeg"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening image file:", err)
	}
	defer file.Close()

	if err != nil {
		fmt.Println("Error reading image file:", err)
	}

	myImage, err := jpeg.Decode(io.Reader(file))

	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer

	if err = jpeg.Encode(&buffer, myImage, nil); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func DownloadImage(url string) ([]byte, error) {
	if url == "" {
		return nil, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, nil
	}

	defer res.Body.Close()

	return JPEGHandler(res)
}
