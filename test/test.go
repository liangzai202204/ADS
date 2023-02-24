package main

import "fmt"

func main() {
	var n int
	var a, b, c int
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	list := make([][]int, n)
	for i := 0; i < n; i++ {
		l := make([]int, 3)
		fmt.Scan(&l[0])
		fmt.Scan(&l[1])
		fmt.Scan(&l[2])
		list = append(list, l)
	}
	fmt.Print("d%\n", list)
}
