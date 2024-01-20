package main

import "fmt"

type Person struct {
	Name    string
	Details *Details
}

type Details struct {
	Age  int
	City string
}

func getPersonCity(p *Person) string {
	return p.Details.City
}

func main() {
	person := &Person{Name: "Simon"}

	city := getPersonCity(person)
	fmt.Println("City:", city)
}
