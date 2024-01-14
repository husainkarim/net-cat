package pkg

import (
	"fmt"
	"log"
)

func PrintError(s string) {
	log.Fatal(s)
}

func PrintUsage() {
	fmt.Println("[USAGE]: ./TCPChat $port")
}
