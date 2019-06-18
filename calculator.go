package main

import (
	"bufio"
	"errors"
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
	var total int64
	var err error
	for index, word := range words {
		if checkIsNumber(word) && index == 0 {
			total, _ = strconv.ParseInt(word, 10, 64)
		} else if index == 0 {
			err = errors.New("First entry must be a number")
		} else {
			switch word {
			case "+":
				total, err = add(total, words[index+1])
			}
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println(total)
}

func checkIsNumber(word string) bool {
	if _, err := strconv.ParseInt(word, 10, 64); err == nil {
		return true
	}
	return false
}

func add(first int64, second string) (int64, error) {
	if checkIsNumber(second) {
		secondNum, _ := strconv.ParseInt(second, 10, 64)
		sum := first + secondNum
		return sum, nil
	}
	return 0, errors.New("not 2 numbers")
}
