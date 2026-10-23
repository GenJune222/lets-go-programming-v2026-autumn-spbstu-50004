package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
	noTemp  = -1
)

type departmentConditioner struct {
	lowerBound int
	upperBound int
}

func newDepartmentConditioner() departmentConditioner {
	return departmentConditioner{lowerBound: minTemp, upperBound: maxTemp}
}

func (c *departmentConditioner) applyEmployeeRequirement(condition string, temperature int) {
	switch condition {
	case ">=":
		c.lowerBound = max(c.lowerBound, temperature)
	case "<=":
		c.upperBound = min(c.upperBound, temperature)
	}
}

func (c *departmentConditioner) findOptimalTemperature() (int, bool) {
	if c.lowerBound > c.upperBound {
		return 0, false
	}
	return c.lowerBound, true
}

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

		conditioner := newDepartmentConditioner()

		for range employeesCount {
			var (
				condition   string
				temperature int
			)

			if _, err := fmt.Scan(&condition, &temperature); err != nil {
				fmt.Println("Invalid temperature")
				return
			}

			conditioner.applyEmployeeRequirement(condition, temperature)

			if optimalTemp, isPossible := conditioner.findOptimalTemperature(); isPossible {
				fmt.Println(optimalTemp)
			} else {
				fmt.Println(noTemp)
			}
		}
	}
}
