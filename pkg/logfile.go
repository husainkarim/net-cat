package pkg

import (
	"bufio"
	"fmt"
	"os"
)

// create log file and add all message
func Logfile(s, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		PrintError("error creating log file")
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, s)
	return writer.Flush()
}
