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
	array := strings.Split(stdin.Text(), "")
	var sum int
	for _, elm := range array {
		v, _ := strconv.Atoi(elm)
		sum += v
	}
	fmt.Printf("%d\n", sum)
	os.Exit(0)
}
