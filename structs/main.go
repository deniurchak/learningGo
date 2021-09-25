package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			zipCode: 12345,
			email: "jim@gmail.com",
		},
	}
	jim.updateName("Carl")	
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName= newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
