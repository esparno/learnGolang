package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}
type SecretAgent struct {
	Person
	LicenseToKill bool
}

func (p Person) FullName() {
	fmt.Println(p.Firstname + " " + p.Lastname)
}

func main() {
	p1 := SecretAgent{}
	p1.Firstname = "James"
	p1.Lastname = "Bond"
	p1.Age = 27
	p1.LicenseToKill = true
	p1.FullName()
}
