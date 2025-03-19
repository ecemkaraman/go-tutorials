```
/go-networking-security
│── main.go
│── server
│   ├── server.go
│── http_client
│   ├── get_post.go
│── json
│   ├── json_parsing.go
│── security
│   ├── hashing.go
│   ├── rsa_keys.go
```

| Topic | Key Concept |
| --- | --- |
| **HTTP Server** | Create & handle GET/POST requests (`net/http`) |
| **HTTP Client** | Send requests (`http.Get()`, `http.Post()`) |
| **JSON Parsing** | Convert struct ↔ JSON (`json.Marshal`, `json.Unmarshal`) |
| **SHA-256 Hashing** | Secure data hashing (`crypto/sha256`) |
| **HMAC Hashing** | HMAC with SHA-256 for authentication (`crypto/hmac`) |
| **RSA Key Generation** | Generate RSA keys for encryption (`crypto/rsa`) |