package main

import "fmt"

func main() {
	var departmentsCount int
	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Println("Invalid departments count")
		return
	}

	for range departmentsCount {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println("Invalid employees count")
			return
		}
	}
}
