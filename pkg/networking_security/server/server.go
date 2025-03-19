package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct for JSON response
type Response struct {
	Message string `json:"message"`
}

// Handle GET request
func handleGet(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "Hello from Go Server!"}
	json.NewEncoder(w).Encode(resp)
}

// Handle POST request
func handlePost(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	json.NewDecoder(r.Body).Decode(&data)
	resp := Response{Message: "Received: " + data["message"]}
	json.NewEncoder(w).Encode(resp)
}

// Start HTTP Server
func StartServer() {
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/post", handlePost)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("❌ Error starting server:", err)
	}
}

/*
🔹 Explanation:
- `http.HandleFunc("/get", handleGet)`: Handles GET requests.
- `json.NewEncoder(w).Encode(resp)`: Sends JSON response.
- `http.ListenAndServe(":8080", nil)`: Starts server on port 8080.
*/
