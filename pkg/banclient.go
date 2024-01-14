package pkg

import (
	"bufio"
	"fmt"
	"os"
)

// add user to ban list
func AddUsrToBanList(usrName string) error {
	var file *os.File
	var err2 error

	_, err := os.Stat("banned-clients.txt")
	if err != nil {
		if os.IsNotExist(err) {
			file, err2 = os.Create("banned-clients.txt")
			if err2 != nil {
				return err2
			}
			defer file.Close()
		} else {
			return err
		}
	}

	file, err2 = os.OpenFile("banned-clients.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err2 != nil {
		return err2
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, usrName)
	return writer.Flush()
}

// open the ban list and read it
func ReadBannedClientName(fName string) ([]string, error) {
	file, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err2 := scanner.Err(); err2 != nil {
		return nil, err2
	}

	return lines, nil
}
