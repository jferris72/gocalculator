package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Let's calculate something \n")
	text, _ := reader.ReadString('\n')
	calculate(strings.Fields(text))
}

func calculate(words []string) {
	var firstNum int64
	for _, word := range words {
		if checkIsNumber(word) {
			firstNum, _ = strconv.ParseInt(word, 10, 64)
			fmt.Print(firstNum)
			fmt.Print("\n")
		}
	}
}

func checkIsNumber(word string) bool {
	if _, err := strconv.ParseInt(word, 10, 64); err == nil {
		return true
	}
	return false
}
