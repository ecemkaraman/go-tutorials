package main

import (
	"fmt"
	"os"

	"github.com/ecemkaraman/go-tutorials/pkg/basics"
	"github.com/ecemkaraman/go-tutorials/pkg/concurrency/channels"
	"github.com/ecemkaraman/go-tutorials/pkg/concurrency/goroutines"
	"github.com/ecemkaraman/go-tutorials/pkg/concurrency/sync"
	"github.com/ecemkaraman/go-tutorials/pkg/data_structures/arrays"
	"github.com/ecemkaraman/go-tutorials/pkg/data_structures/maps"
	"github.com/ecemkaraman/go-tutorials/pkg/data_structures/pointers"
	"github.com/ecemkaraman/go-tutorials/pkg/data_structures/slices"
	"github.com/ecemkaraman/go-tutorials/pkg/data_structures/structs"
	errorhandling "github.com/ecemkaraman/go-tutorials/pkg/error_handling"
	filehandling "github.com/ecemkaraman/go-tutorials/pkg/file_handling"
	"github.com/ecemkaraman/go-tutorials/pkg/networking_security/http_client"
	"github.com/ecemkaraman/go-tutorials/pkg/networking_security/json"
	"github.com/ecemkaraman/go-tutorials/pkg/networking_security/security"
	"github.com/ecemkaraman/go-tutorials/pkg/networking_security/server"
	"github.com/ecemkaraman/go-tutorials/pkg/practical_apps/api"
	"github.com/ecemkaraman/go-tutorials/pkg/practical_apps/caching"
	"github.com/ecemkaraman/go-tutorials/pkg/practical_apps/cli"
	"github.com/ecemkaraman/go-tutorials/pkg/practical_apps/scraper"
	"github.com/ecemkaraman/go-tutorials/pkg/practical_apps/worker"
)

// This is for future use of entry points into reusable packages for Go
func main() {
	fmt.Println("Select a topic to run:")
	fmt.Println("0 - Go Basics")
	fmt.Println("1 - Arrays")
	fmt.Println("2 - Maps")
	fmt.Println("3 - Slices")
	fmt.Println("4 - Structs")
	fmt.Println("5 - Pointers")
	fmt.Println("6 - Goroutines")
	fmt.Println("7 - Channels")
	fmt.Println("8 - Mutex & Sync")
	fmt.Println("9 - Errors Handling")
	fmt.Println("10 - File Handling")
	fmt.Println("11 - HTTP Client")
	fmt.Println("12 - JSON Parsing")
	fmt.Println("13 - Security")
	fmt.Println("14 - HTTP Server")
	fmt.Println("15 - REST API")
	fmt.Println("16 - Caching")
	fmt.Println("17 - CLI Tool")
	fmt.Println("18 - Web Scraper")
	fmt.Println("19 - Worker Pool")

	var choice int
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input:", err)
		os.Exit(1)
	}

	switch choice {
	case 0:
		basics.Run()
	case 1:
		arrays.Run()
	case 2:
		maps.Run()
	case 3:
		slices.Run()
	case 4:
		structs.Run()
	case 5:
		pointers.Run()
	case 6:
		goroutines.Run()
	case 7:
		channels.Run()
	case 8:
		sync.Run()
	case 9:
		errorhandling.Run()
	case 10:
		filehandling.Run()
	case 11:
		http_client.Run()
	case 12:
		json.Run()
	case 13:
		security.Run()
	case 14:
		server.Run()
	case 15:
		api.Run()
	case 16:
		caching.Run()
	case 17:
		cli.Run()
	case 18:
		scraper.Run()
	case 19:
		worker.Run()
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
	}
}
