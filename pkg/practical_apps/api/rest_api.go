package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ✅ Handler function for GET request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, Go API!"}
	json.NewEncoder(w).Encode(response)
}

// ✅ Start an HTTP server
func Run() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

/*
🔹 Explanation:
- `http.HandleFunc()`: Registers the handler function.
- `http.ListenAndServe()`: Starts server at `localhost:8080`.
- `curl http://localhost:8080/hello`
*/
