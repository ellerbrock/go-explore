package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := sum(1, 2, 3, 4, 5)
	fmt.Println("\nresult:", res)

	counter(10)

}

func sum(vals ...int) int {
	sum := 0
	for i := range vals {
		fmt.Println("add: ", vals[i])
		sum += vals[i]
		fmt.Println("sum:" + strconv.Itoa(sum))
	}
	return sum
}

func counter(cnt int) {
	for i := cnt; i >= 0; i-- {
		fmt.Printf("Counter: %d\n", i)
	}

}
