package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}

func main() {
	type person struct {
		firstName string
		lastName  string
		sex       string
		company   string
	}

	jasiu := person{
		firstName: "Jan",
		lastName:  "Kowalski",
		sex:       "Yes, please",
		company:   "Samsung",
	}

	fmt.Printf("Name: %s %s Sex: %s Company: %s\n",
		jasiu.firstName, jasiu.lastName,
		jasiu.sex, jasiu.company,
	)
	arr := []int64{1, 2, 3, 10, 20}
	fmt.Println(arr)
	arr2 := []int64{10, 10, 20}
	arr = append(arr, arr2...)
}
