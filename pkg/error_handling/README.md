```
/go-error-file-handling
│── main.go
│── errorhandling
│   ├── errors.go
│   ├── panic_recover.go
│   ├── custom_errors.go
│── filehandling
│   ├── read_file.go
│   ├── write_file.go
│   ├── append_file.go
```

### **🚀 Summary**

| Topic | Key Concept |
| --- | --- |
| **Basic Errors** | Create (`errors.New`), format (`fmt.Errorf`), check (`if err != nil`) |
| **Panic & Recover** | Trigger (`panic()`), handle (`recover()`) |
| **Custom Errors** | Define (`struct`), implement (`Error() string`) |
| **Read File** | `ioutil.ReadFile("file")` |
| **Write File** | `ioutil.WriteFile("file", []byte, perm)` |
| **Append File** | `os.OpenFile("file", os.O_APPEND|os.O_WRONLY, 0644)` |