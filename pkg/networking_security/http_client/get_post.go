package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Runs HTTP GET & POST requests
func RunHTTPClient() {
	// ✅ GET Request
	response, err := http.Get("http://localhost:8080/get")
	if err != nil {
		fmt.Println("❌ GET Request Error:", err)
		return
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("✅ GET Response:", string(body))

	// ✅ POST Request
	data := map[string]string{"message": "Hello, Server!"}
	jsonData, _ := json.Marshal(data)
	resp, err := http.Post("http://localhost:8080/post", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("❌ POST Request Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println("✅ POST Response:", string(body))
}

/*
🔹 Explanation:
- `http.Get("url")`: Sends GET request.
- `http.Post("url", "content-type", body)`: Sends POST request.
- `json.Marshal(data)`: Converts struct/map to JSON.
*/
