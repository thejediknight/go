package main

import (
	"fmt"
	"errors"
)

var _p = fmt.Println
var _sp = fmt.Sprintf

//~ Define a new type called User which is a struct.
type User struct {
	Name, Role, Email string
	Age int
}

//~ Extension method for type User. 
//# Method name is Salary() taking 0 params and returns int.
func (u User) Salary() (int, error) {
	switch u.Role {
	case "Developer":
		return 100, nil
	case "Architect":
		return 200, nil
	case "Clerk":
		return 50, nil
	default:
		return 0, errors.New(
			_sp("Error! Role not found for %s!", u.Role))
	}
}

//~ Extension method for type User. 
//# Method name is UpdateEmail() taking 1 param and returns void.
func (u *User) UpdateEmail(email string) {
	u.Email = email
}

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

func mapGo() {
	price := map[string]int {
		"apple": 12,
		"banana": 10,
		"peach": 9,
	} 

	//~ Print the full map
	_p(price)

	//~ Print value of a specified key. Check for key existance as well.
	//# Define a slice of keys. Loop through each key and print the value if it exists in map.
	keys := []string{"apple", "pineapple"}
	for _, key := range keys {
		value, exists := price[key]
		if exists == true {
			_p(key, "=>", value)
		} else {
			_p(key, "doesn't exist in map")
		}
	
		//~ Print using range
		for index, value := range price {
			if index != "banana" {
				_p(index, "=>", value)
			}
		}	
	}

	//~ Add another entry in map
	price["Avocado"] = 100
	_p(price)
}

func structGo() {
	// Create a new user from the type. Can omit a field.
	user1 := User {
		Name: "Bose Dk",
		Role: "Developer",
		Age: 18,
	}

	// Print the struct
	_p(user1)

	// Print salary of this user
	_p(user1.Salary())

	// Update the email of this user.
	user1.UpdateEmail("bosedk@gmail.com")

	// Print the struct
	_p(user1)

	// Create a new user with unsupported Role, and try to print its salary. We expect custom error.
	user2 := User{
		Name: "Mk",
		Role: "Joker",
	}
	
	if salary, err := user2.Salary(); err != nil {
		_p(err)
	} else {
		_p(salary)
	}
}