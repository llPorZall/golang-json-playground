package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/llporzall/golang-json-playground/user"
)

func commandParser(cmd string) {
	inputArg := strings.Fields(cmd)
	if len(inputArg) == 1 || len(inputArg) == 3 {
		if inputArg[0] == "ADD" {
			age, err := strconv.Atoi(inputArg[2])
			if err != nil {
				fmt.Println("Input age only integer")
			} else {
				user.AddNewUser(inputArg[1], age)
			}
		} else if inputArg[0] == "READ" {
			user.GetAllUser()
		} else {
			fmt.Println("Command not found,Please endter new command again")
		}
	} else {
		fmt.Println("Command not found,Please endter new command again")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please Enter command \nfor the example")
	fmt.Println("ADD PONGCHAI 43")
	fmt.Println("Or exit for exit app")
	for {
		scanner.Scan()
		text := scanner.Text()
		if text == "exit" || len(text) == 0 {
			break
		} else {
			commandParser(text)
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}
