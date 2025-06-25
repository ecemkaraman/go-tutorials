package security

import "fmt"

func Run() {
	fmt.Println("=== SHA-256 Hash Demo ===")
	RunSHA256Hash()

	fmt.Println("\n=== HMAC Hash Demo ===")
	RunHMACHash()

	fmt.Println("\n=== RSA Key Generation Demo ===")
	RunRSAKeyGen()
}
