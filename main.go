package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		byts, _, e := reader.ReadLine()
		byt2, _, _ := reader.ReadLine()
		input := string(byts)
		nums := string(byt2)
		if e != nil || input == "" {
			break
		}
		num, _ := strconv.Atoi(nums)
		fmt.Println(input[:num])
	}
}
