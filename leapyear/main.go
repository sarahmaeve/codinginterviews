package main

import "fmt"

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%400 == 0 {
			return true
		} else if year%100 == 0 {
			fmt.Println("Divis by 100")
			return false

		} else {
			return true
		}
	} else {
		fmt.Println("Not divis by four")
		return false

	}
}

func main() {

	fmt.Println(isLeapYear(2000))
	fmt.Println(isLeapYear(2012))
	fmt.Println(isLeapYear(1900)) // false
	fmt.Println(isLeapYear(1919))
}
