package main

import "fmt"

var _p = fmt.Println

func swap(x, y int) (int, int) {
	return y, x
}

func swapMain() {
	x := 10
	y := 13

	x,y = swap(x,y)

	fmt.Println(x, y)
}

func loop() {
	for i:=0; i<10; i++ {
		_p(i)
	}
}

func price(item string) int {
	switch item {
	case "apple", "tomatoes":
		return 10
	case "peach":
		return 5
	case "banana":
		return 9
	default:
		return -1
	}
}

func main() {
	_p(price("apple"))
}