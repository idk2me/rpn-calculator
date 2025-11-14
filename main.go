package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"example.com/m/calculator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter an expression in RPN: ")
	expr, _ := reader.ReadString('\n')
	expr = strings.TrimSpace(expr)

	res, err := calculator.Calc(expr)
	if err != nil {
		log.Fatalf("An error has occured, please check your expression %s: %s", expr, err)
	}
	fmt.Printf("Ans: %v", res)
}
