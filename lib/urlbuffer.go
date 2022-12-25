package lib

import (
	"io"
	"log"
	"net/http"
	"os"
)

func GetBuffer(URL, fileName string) string {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatal(response.StatusCode)
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return fileName
}