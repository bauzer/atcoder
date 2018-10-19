package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	s := strings.Split(stdin.Text(), " ")
	first, _ := strconv.Atoi(s[0])
	second, _ := strconv.Atoi(s[1])

	var answer string
	if first%2 == 1 && second%2 == 1 {
		answer = "Odd"
	} else {
		answer = "Even"
	}
	fmt.Printf("%s\n", answer)
	os.Exit(0)
}
