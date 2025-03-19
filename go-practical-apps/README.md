```
/go-practical-apps
│── main.go                   # Entry point of the project
│── go.mod                    # Module file for dependencies
│── go.sum                    # Checksum for dependencies
│── README.md                 # Project documentation
│
│── cli/                      # 🛠 CLI Tool
│   ├── cli_tool.go           # CLI implementation using flag package
│
│── api/                      # 🌐 REST API
│   ├── rest_api.go           # REST API implementation with net/http
│
│── scraper/                  # 🕸 Web Scraper
│   ├── web_scraper.go        # Scrapes web pages for links
│
│── caching/                  # 💾 Caching Mechanism
│   ├── cache.go              # In-memory cache with sync.RWMutex
│
│── worker/                   # ⚡ Worker Pool
│   ├── worker_pool.go        # Concurrency with Goroutines & Channels

```

### **🚀 Summary**

| Topic | Key Concept |
| --- | --- |
| **CLI Tool** | `flag.String()`, `flag.Int()`, `flag.Parse()` |
| **REST API** | `http.HandleFunc()`, `http.ListenAndServe()` |
| **Web Scraper** | `golang.org/x/net/html`, `html.NewTokenizer()` |
| **Caching** | `map`, `sync.RWMutex` |
| **Worker Pool** | `goroutines`, `channels`, `sync.WaitGroup` |