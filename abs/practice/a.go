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
	first, _ := strconv.Atoi(stdin.Text())
	stdin.Scan()
	s := strings.Split(stdin.Text(), " ")
	second, _ := strconv.Atoi(s[0])
	third, _ := strconv.Atoi(s[1])
	stdin.Scan()
	text := stdin.Text()

	answer := first + second + third
	fmt.Printf("%d %s\n", answer, text)
	os.Exit(0)
}
