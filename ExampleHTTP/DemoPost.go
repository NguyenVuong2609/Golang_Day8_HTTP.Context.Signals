package ExampleHTTP

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func TestHttpPost() {
	url := "https://jsonplaceholder.typicode.com/todos"

	// Create data then parse into []byte
	jsonData := Todo{10, "Hello", 201, true}
	parseData, err := json.Marshal(jsonData)

	// Create Post Request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(parseData))
	if err != nil {
		log.Fatal(err)
	}

	// Set up Header:  Content-Type application/json
	request.Header.Set("Content-Type", "application/json")

	// Send request and get response
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Show Http status code:
	fmt.Println("HTTP Status Code:", response.StatusCode)
}
