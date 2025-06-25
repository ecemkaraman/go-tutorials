package filehandling

import "fmt"

func Run() {
	fmt.Println("=== Reading File ===")
	RunReadFile()

	fmt.Println("\n=== Writing File ===")
	RunWriteFile()

	fmt.Println("\n=== Appending to File ===")
	RunAppendFile()
}
