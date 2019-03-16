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

func even() {
	for i:=0; i<10; i++ {
		if i%2 == 0 {
			_p(i)
		}
	}
}

func switchFunc() {
	closure := func (item string) int {
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

	_p(closure("apple"))
}

func sliceFunc() {
	//~ Init a slice of int.
	evenNums := []int{2,4,6,8}
	
	//~ Regular loop
	for i:=0; i<len(evenNums); i++ {
		_p(evenNums[i])
	}

	//~ Loop using range. Range returns (index, value)
	for index, value := range evenNums {
		_p(index, "=>", value)
	} 

	//~ Ignore index, just print value
	for _, value := range evenNums {
		_p("Value:", value)
	} 

	//~ Ignore value, just print index
	for index, _ := range evenNums {
		_p("Index:", index)
	} 
	
	//~ Just print the list of numbers in evenNums [2 4 6 8]
	_p(evenNums)

	//~ Append 100, 101, 102 to the list evenNums. It doesnt modify evenNums.
	list2 := append(evenNums, 100, 101, 102)
	_p(evenNums, list2)

}

func main() {
	sliceFunc()
}