package main

import (
	"fmt"
	"os"

	"github.com/ecemkaraman/go-tutorials/pkg/concurrency/channels"
	"github.com/ecemkaraman/go-tutorials/pkg/concurrency/goroutines"
	"github.com/ecemkaraman/go-tutorials/pkg/concurrency/sync"
	"github.com/ecemkaraman/go-tutorials/pkg/data_structures/arrays"
	"github.com/ecemkaraman/go-tutorials/pkg/data_structures/maps"
	"github.com/ecemkaraman/go-tutorials/pkg/data_structures/pointers"
	"github.com/ecemkaraman/go-tutorials/pkg/data_structures/slices"
	"github.com/ecemkaraman/go-tutorials/pkg/data_structures/structs"
	filehandling "github.com/ecemkaraman/go-tutorials/pkg/file_handling/append_file"
	filehandling "github.com/ecemkaraman/go-tutorials/pkg/file_handling/read_file"
	filehandling "github.com/ecemkaraman/go-tutorials/pkg/file_handling/write_file"
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

func main() {
	fmt.Println("Select a topic to run:")
	fmt.Println("1 - Arrays")
	fmt.Println("2 - Maps")
	fmt.Println("3 - Slices")
	fmt.Println("4 - Structs")
	fmt.Println("5 - Pointers")
	fmt.Println("6 - Goroutines")
	fmt.Println("7 - Channels")
	fmt.Println("8 - Mutex & Sync")
	fmt.Println("9 - Errors Handling")
	fmt.Println("10 - Panic & Recover")
	fmt.Println("11 - Read File")
	fmt.Println("12 - Write File")
	fmt.Println("13 - HTTP Client")
	fmt.Println("14 - JSON Parsing")
	fmt.Println("15 - Security")
	fmt.Println("16 - HTTP Server")
	fmt.Println("17 - REST API")
	fmt.Println("18 - Caching")
	fmt.Println("19 - CLI Tool")
	fmt.Println("20 - Web Scraper")
	fmt.Println("21 - Worker Pool")

	var choice int
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input:", err)
		os.Exit(1)
	}

	switch choice {
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
		channels.RunBufferedChannel()
		channels.RunUnbufferedChannel()
	case 8:
		sync.RunMutex()
		sync.RunRWMutex()
		sync.RunWaitGroup()
	case 9:
		errorhandling.RunErrors()
		errorhandling.RunPanic()
	case 10:
		errorhandling.RunPanic()
	case 11:
		filehandling.RunReadFile()
	case 12:
		filehandling.RunWriteFile()
	case 13:
		http_client.Run()
	case 14:
		json.Run()
	case 15:
		security.RunSHA256Hash()
		security.RunHMACHash()
		security.RunRSAKeyGen()
	case 16:
		server.StartServer()
	case 17:
		api.Run()
	case 18:
		caching.Run()
	case 19:
		cli.Run()
	case 20:
		scraper.Run()
	case 21:
		worker.Run()
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
	}
}
