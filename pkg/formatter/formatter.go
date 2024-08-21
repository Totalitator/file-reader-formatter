package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Formating(f Formatter) string {
	fmt.Println("Formatted text:", f.Format())
	return f.Format()
}

func GetFormatter(method string, data string) Formatter {
	switch method {
	case methodPrint:
		return &Print{Text: data}
	case methodUpperCase:
		return &UpperCase{Text: data}
	case methodLowerCase:
		return &LowerCase{Text: data}
	default:
		log.Fatalf("Unknown method: %s", method)
		return nil
	}
}

func SaveFormat(file *os.File, newFormat string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Wanna save format? yes/no?...")
	input, _ := reader.ReadString('\n')
	if strings.TrimSpace(input) == "yes" {
		file, err := os.OpenFile(file.Name(), os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err = file.WriteString(newFormat + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
