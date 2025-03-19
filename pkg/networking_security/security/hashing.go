package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// Runs SHA-256 & HMAC Hashing
func RunSHA256Hash() {
	data := "secure data"
	hash := sha256.Sum256([]byte(data))
	fmt.Printf("✅ SHA-256 Hash: %x\n", hash)
}

func RunHMACHash() {
	key := []byte("secret-key")
	data := []byte("secure data")
	h := hmac.New(sha256.New, key)
	h.Write(data)
	fmt.Printf("✅ HMAC Hash: %x\n", h.Sum(nil))
}

/*
🔹 Explanation:
- `sha256.Sum256([]byte(data))`: Hashes data using SHA-256.
- `hmac.New(sha256.New, key)`: Creates HMAC with SHA-256.
*/
