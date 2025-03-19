# go-tutorials
Playground for Golang


## **📂 Go Tutorials Repository Structure**

A **comprehensive Go repository** covering **core language features, concurrency, networking, security, interfaces, error handling, advanced topics, and practical applications.**
```
/go-tutorials
│── go.mod                            # Go module file
│── go.sum                            # Dependency checksum
│── README.md                         # Main project documentation
│
│── go-core/                           # 🟢 Go Core Features
│   ├── variables.go                   # Variables & Data Types
│   ├── operators.go                   # Operators & Expressions
│   ├── control_flow.go                 # If, Switch, Loops
│   ├── functions.go                    # Functions & Recursion
│   ├── README.md                       # Documentation
│
│── go-data-structures/                # 🔹 Data Structures
│   ├── arrays.go                       # Arrays
│   ├── slices.go                       # Slices
│   ├── maps.go                         # Maps
│   ├── structs.go                      # Structs
│   ├── pointers.go                     # Pointers
│   ├── README.md                       # Documentation
│
│── go-concurrency/                    # ⚡ Concurrency & Goroutines
│   ├── goroutines.go                   # Goroutines Basics
│   ├── channels.go                     # Channels
│   ├── buffered_channels.go            # Buffered vs Unbuffered
│   ├── waitgroups.go                   # sync.WaitGroup
│   ├── mutex.go                         # sync.Mutex & RWMutex
│   ├── README.md                        # Documentation
│
│── go-networking-security/            # 🔒 Networking & Security
│   ├── http_server.go                   # HTTP Server
│   ├── http_client.go                   # GET & POST Requests
│   ├── json_parsing.go                  # JSON Parsing
│   ├── hashing.go                        # SHA-256 & HMAC
│   ├── rsa_keys.go                       # RSA Key Generation
│   ├── README.md                         # Documentation
│
│── go-interfaces-polymorphism/        # 🎭 Interfaces & Polymorphism
│   ├── interfaces.go                    # Defining & Implementing Interfaces
│   ├── polymorphism.go                   # Interface-based Polymorphism
│   ├── type_assertions.go                # Type Assertions & Switches
│   ├── README.md                         # Documentation
│
│── go-error-file-handling/            # 🚨 Error & File Handling
│   ├── errors.go                        # errors.New() & fmt.Errorf()
│   ├── panic_recover.go                 # Panic & Recover
│   ├── custom_errors.go                 # Custom Error Types
│   ├── file_reading.go                  # os.Open(), ioutil.ReadFile()
│   ├── file_writing.go                  # os.Create(), ioutil.WriteFile()
│   ├── file_appending.go                # os.OpenFile() Append Mode
│   ├── README.md                         # Documentation
│
│── go-advanced/                        # 🔥 Advanced Go Topics
│   ├── reflection.go                    # Using the reflect package
│   ├── context.go                        # Managing Goroutines with Context
│   ├── sync_pool.go                      # Memory Optimization with sync.Pool
│   ├── benchmarking.go                   # Benchmarking with testing.B
│   ├── modules.go                        # Creating & Using Modules
│   ├── README.md                         # Documentation
│
│── go-practical-apps/                   # 🛠 Practical Go Applications
│   ├── cli_tool.go                       # Building CLI Tool with flag & cobra
│   ├── rest_api.go                        # Creating a REST API
│   ├── web_scraper.go                     # Web Scraper using net/html
│   ├── caching.go                         # Caching Mechanism with map & sync.RWMutex
│   ├── worker_pool.go                      # Worker Pool with Goroutines & Channels
│   ├── README.md                           # Documentation
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