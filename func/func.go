package func

import "fmt"

func add(x int, y int) int {
	return x + y
}

func Do() {
	fmt.Println(add(42, 13))
}
