package pkg

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateFile() *os.File {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Name your file...")

	filename, _ := reader.ReadString('\n')

	filename = strings.TrimSpace(filename)

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(filename + ".txt")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		return file
	} else {
		fmt.Printf("File with name '%s' already exists, recreate the file with new, unused name", filename)
		return nil
	}
}

func FillFile(file *os.File) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Fill in the file with the information...")

	input, _ := reader.ReadString('\n')

	_, err := file.WriteString(strings.TrimSpace(input))
	if err != nil {
		log.Fatal(err)
	}
}

func Reader(file *os.File) string {
	bytes, err := os.ReadFile(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	data := string(bytes)
	return data
}
