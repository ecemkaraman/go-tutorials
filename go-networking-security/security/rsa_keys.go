package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// Generates RSA Private & Public Keys
func RunRSAKeyGen() {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	// ✅ Save Private Key
	privFile, _ := os.Create("private.pem")
	pem.Encode(privFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	privFile.Close()
	fmt.Println("✅ RSA Private Key Generated: private.pem")

	// ✅ Save Public Key
	publicKey := &privateKey.PublicKey
	pubFile, _ := os.Create("public.pem")
	pem.Encode(pubFile, &pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(publicKey)})
	pubFile.Close()
	fmt.Println("✅ RSA Public Key Generated: public.pem")
}

/*
🔹 Explanation:
- `rsa.GenerateKey(rand.Reader, 2048)`: Generates 2048-bit RSA key.
- `x509.MarshalPKCS1PrivateKey(privateKey)`: Encodes private key.
- `pem.Encode(file, &pem.Block{Type, Bytes})`: Saves keys to file.
*/
