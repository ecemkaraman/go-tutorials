# go-tutorials
Playground for Golang


## **📂 Go Tutorials Repository Structure**

A **comprehensive Go repository** covering **core language features, concurrency, networking, security, interfaces, error handling, advanced topics, and practical applications.**
```
go-tutorials/
├── cmd/                                 # 🚀 Central Runner
│   └── main.go                          # CLI-based entrypoint to explore all modules interactively

├── go.mod                               # 📦 Go module file

├── pkg/                                 # 🧱 Main Go Codebase
│
│── advanced/                            # 🔥 Advanced Go Topics
│   ├── reflection/                      # Using the reflect package
│   │   └── reflection.go
│   ├── concurrency/                     # Managing Goroutines with context
│   │   └── context.go
│   ├── memory/                          # Memory optimization strategies
│   │   └── sync_pool.go
│   ├── modules/                         # Creating and using Go modules
│   │   └── go_modules.go
│   ├── testing/                         # Benchmarking with testing.B
│   │   └── benchmark_test.go
│   └── run.go                           # Entrypoint for advanced topics
│
│── concurrency/                         # ⚡ Concurrency & Parallelism
│   ├── goroutines/                      # Goroutine basics and launching
│   │   └── goroutines.go
│   ├── channels/                        # Communication through channels
│   │   ├── buffered.go                  # Buffered channel example
│   │   ├── unbuffered.go                # Unbuffered channel example
│   │   └── run.go                       # Unified entrypoint for channels
│   ├── sync/                            # sync tools for coordination
│   │   ├── mutex.go                     # sync.Mutex demo
│   │   ├── rwmutex.go                   # sync.RWMutex demo
│   │   ├── waitgroup.go                 # sync.WaitGroup demo
│   │   └── run.go                       # Unified entrypoint for sync
│
│── control_flow/                        # 🔄 Control Structures
│   └── controlflow.go                   # If, switch, and loops demo
│
│── data_structures/                    # 🧊 Core Data Structures
│   ├── arrays/                          # Fixed-size indexed storage
│   │   └── arrays.go
│   ├── slices/                          # Dynamic, resizable arrays
│   │   └── slices.go
│   ├── maps/                            # Key-value store
│   │   └── maps.go
│   ├── structs/                         # Custom types and methods
│   │   └── structs.go
│   ├── pointers/                        # Working with memory addresses
│   │   └── pointers.go
│   ├── strings/                         # String operations and tricks
│   │   └── strings.go
│   └── run.go                           # Entrypoint to run all data structure demos
│
│── error_handling/                     # 🚨 Error Handling in Go
│   ├── standard_errors.go               # Basic errors.New and fmt.Errorf
│   ├── custom_errors.go                 # Defining custom error types
│   ├── panic.go                         # Panic & recover mechanisms
│   └── run.go                           # Unified error handling entrypoint
│
│── file_handling/                      # 📁 Working with Files
│   ├── read_file.go                     # Reading with os and ioutil
│   ├── write_file.go                    # Writing to files
│   ├── append_file.go                   # Appending to files using os.OpenFile
│   └── run.go                           # Entry point to run all file operations
│
│── interfaces/                         # 🎭 Interfaces & Polymorphism
│   ├── basic_interface.go               # Defining and using interfaces
│   ├── interface_as_param.go            # Passing interfaces as function parameters
│   ├── empty_interface.go               # `interface{}` and type safety
│   ├── type_assertions.go               # Safe type assertions
│   ├── type_switch.go                   # Handling multiple types via switch
│   └── run.go                           # Entrypoint for interface demos
│
│── networking_security/               # 🔒 Networking & Security
│   ├── http_client/                    # HTTP GET/POST requests
│   │   └── get_post.go
│   ├── json/                           # Working with JSON
│   │   └── json_parsing.go
│   ├── security/                       # Cryptography tools
│   │   ├── hashing.go                  # SHA-256, HMAC
│   │   ├── rsa_keys.go                 # RSA keypair generation
│   │   └── run.go                      # Unified crypto demo entrypoint
│   ├── server/                         # Serving content over HTTP
│   │   └── server.go
│
│── practical_apps/                    # 🛠 Practical Applications
│   ├── cli/                            # Building CLI tools with `flag`
│   │   └── cli_tool.go
│   ├── api/                            # Building a REST API with `net/http`
│   │   └── rest_api.go
│   ├── scraper/                        # Basic web scraper using `net/html`
│   │   └── web_scraper.go
│   ├── caching/                        # In-memory cache with sync.RWMutex
│   │   └── cache.go
│   ├── worker/                         # Worker pool pattern with channels
│   │   └── worker_pool.go
│   └── run.go                          # Entrypoint to run practical app demos
```




## **📌 How to Use**

### **🔹 1. Install Go**

Ensure you have **Go installed** on your system.

🔗 [**Download & Install Go**](https://go.dev/dl/)

Verify installation:

```
go version

```

### **🔹 2. Clone the Repository**

```
git clone <https://github.com/ecemkaraman/go-tutorials.git>
cd go-tutorials

```

### **🔹 3. Run Each Section Individually**

You can **navigate into each folder** and run specific Go programs.

### ✅ **Run Concurrency Examples**

```
cd go-concurrency
go run goroutines.go

```

### ✅ **Run Practical Applications**

```
cd go-practical-apps
go run cli_tool.go

```

---

## **📌 Key Features**

| Section | Description |
| --- | --- |
| 🟢 **Go Core** | Variables, Data Types, Control Flow, Functions |
| 🔹 **Data Structures** | Arrays, Slices, Maps, Structs, Pointers |
| ⚡ **Concurrency** | Goroutines, Channels, Mutexes, WaitGroups |
| 🔒 **Networking & Security** | HTTP Server, JSON Parsing, Hashing, RSA |
| 🎭 **Interfaces & Polymorphism** | Implementing Interfaces, Type Assertions |
| 🚨 **Error & File Handling** | Errors, Panic/Recover, File Reading & Writing |
| 🔥 **Advanced Go** | Reflection, Context, Memory Optimization, Modules |
| 🛠 **Practical Apps** | CLI Tools, REST API, Web Scraper, Caching, Worker Pool |

---

## **📌 Dependencies**

This project uses Go modules. Install dependencies with:

```
go mod tidy

```

---

## **📌 Contributing**

Want to improve this project?

1. **Fork the repository**
2. **Create a new branch** (`feature-new` or `fix-bug`)
3. **Commit your changes**
4. **Push & create a pull request**

---

## **📌 License**

MIT License © [Ecem Karaman](https://github.com/ecemkaraman)
