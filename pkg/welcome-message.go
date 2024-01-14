package pkg

import (
	"io"
	"os"
)

// welcome message
func OpenTxt(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		PrintError("error opening txt file")
		return []byte{}
	}
	defer file.Close()
	result, _ := io.ReadAll(file)
	return result
}

func WelcomeMessage() []byte {
	welmes := OpenTxt("linux-logo.txt")
	return welmes
}
