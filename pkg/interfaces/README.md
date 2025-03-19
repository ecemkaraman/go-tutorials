```
/go-interfaces-polymorphism
│── main.go
│── interfaces
│   ├── basic_interface.go
│   ├── interface_as_param.go
│   ├── empty_interface.go
│   ├── type_assertions.go
│   ├── type_switch.go
```

### **🚀 Summary**

| Topic | Key Concept |
| --- | --- |
| **Basic Interface** | `type InterfaceName interface {}` |
| **Passing Interface** | `func Speak(a Animal)` |
| **Empty Interface** | `func Anything(v interface{})` (accepts all types) |
| **Type Assertions** | `value, ok := i.(Type)` |
| **Type Switch** | `switch v := i.(type) { case int: ... }` |