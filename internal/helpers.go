package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ChooseMethod() string {
	var method string
	fmt.Println("Print any methods: print/uppercase/lowercase:   ")
	fmt.Scanln(&method)

	return method
}

func Cycle() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Wanna try again? yes/no?...")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input) == "yes"

}
