package main

import (
	"fmt"

	"github.com/ecemkaraman/go-tutorials/go-networking-security/http_client"
	"github.com/ecemkaraman/go-tutorials/go-networking-security/json"
	"github.com/ecemkaraman/go-tutorials/go-networking-security/security"
	"github.com/ecemkaraman/go-tutorials/go-networking-security/server"
)

func main() {
	fmt.Println("🚀 Go Networking & Security Practice")

	// ✅ Start HTTP Server
	go server.StartServer()
	fmt.Println("🔹 HTTP Server Started on :8080")

	// ✅ Run HTTP Client
	fmt.Println("\n🔹 Running HTTP Client Example")
	http_client.RunHTTPClient()

	// ✅ Run JSON Parsing
	fmt.Println("\n🔹 Running JSON Parsing Example")
	json.RunJSONParsing()

	// ✅ Run SHA-256 Hashing
	fmt.Println("\n🔹 Running SHA-256 Hashing Example")
	security.RunSHA256Hash()

	// ✅ Run HMAC Hashing
	fmt.Println("\n🔹 Running HMAC Hashing Example")
	security.RunHMACHash()

	// ✅ Run RSA Key Generation
	fmt.Println("\n🔹 Running RSA Key Generation Example")
	security.RunRSAKeyGen()

	// Prevent main from exiting immediately
	select {}
}
