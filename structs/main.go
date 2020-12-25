package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	saeed := person{
		firstName: "saeed",
		lastName:  "kazemi",
		contact: contactInfo{
			email:   "saeed@gmail.com",
			zipCode: 0513,
		},
	}
	saeed.updateName("ferydoon")
	saeed.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
