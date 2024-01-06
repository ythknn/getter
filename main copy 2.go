package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// URL base for the requests
	baseURL := "http://kulturenvanteri.com/tr/wp-json/ke/place/"

	file, err := os.Create("responses2.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for i := 600001; i <= 1000000; i++ {

		url := baseURL + strconv.Itoa(i)

		// Make the HTTP GET request
		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making request for", url, ":", err)
			continue
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body for", url, ":", err)
			continue
		}

		fmt.Println("Response for", url, ":", string(body))

		_, err = file.WriteString(fmt.Sprintf("Response for %s: %s\n", url, string(body)))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Requests completed. Responses saved to responses.txt")
}
