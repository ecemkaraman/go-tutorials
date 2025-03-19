package main

import (
	"fmt"

	"github.com/ecemkaraman/go-tutorials/go-practical-apps/api"
	"github.com/ecemkaraman/go-tutorials/go-practical-apps/caching"
	"github.com/ecemkaraman/go-tutorials/go-practical-apps/cli"
	"github.com/ecemkaraman/go-tutorials/go-practical-apps/scraper"
	"github.com/ecemkaraman/go-tutorials/go-practical-apps/worker"
)

func main() {
	fmt.Println("🚀 Go Practical Applications")

	// ✅ CLI Tool
	fmt.Println("\n🔹 Running CLI Tool")
	cli.RunCLI()

	// ✅ REST API
	fmt.Println("\n🔹 Running REST API")
	go api.StartServer()

	// ✅ Web Scraper
	fmt.Println("\n🔹 Running Web Scraper")
	scraper.RunScraper()

	// ✅ Caching Mechanism
	fmt.Println("\n🔹 Running Caching Mechanism")
	caching.RunCache()

	// ✅ Worker Pool
	fmt.Println("\n🔹 Running Worker Pool")
	worker.RunWorkerPool()
}
