package json

import (
	"encoding/json"
	"fmt"
)

// Struct for JSON data
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Runs JSON Encoding & Decoding Example
func RunJSONParsing() {
	// ✅ Encode Struct to JSON
	user := User{Name: "Ecem", Email: "ecem@example.com"}
	jsonData, _ := json.Marshal(user)
	fmt.Println("✅ JSON Encoded:", string(jsonData))

	// ✅ Decode JSON to Struct
	var decodedUser User
	json.Unmarshal(jsonData, &decodedUser)
	fmt.Println("✅ JSON Decoded:", decodedUser)
}

/*
🔹 Explanation:
- `json.Marshal(data)`: Converts struct → JSON.
- `json.Unmarshal(jsonData, &struct)`: Converts JSON → struct.
*/
