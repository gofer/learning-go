package main

import "fmt"

func exercise03() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	employee1 := Employee{"田中", "太郎", 123}

	employee2 := Employee{
		firstName: "鈴木",
		lastName:  "花子",
		id:        456,
	}

	var employee3 Employee
	employee3.firstName = "山田"
	employee3.lastName = "一郎"
	employee3.id = 789

	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println(employee3)
}
