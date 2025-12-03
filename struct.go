package main

import "fmt"

func main() {
	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       42,
		address: Address{
			city:    "Jakarta",
			country: "Indonesia",
		},
		HomePhoneCell: HomePhoneCell{
			home: "Tj. Duren",
			cell: "08123123",
		},
	}

	fmt.Println(p.firstName)
	fmt.Println(p.fullName())

	fmt.Println("Before increment", p.age)
	p.incrementAgeByOne()
	fmt.Println("After increment", p.age)
	fmt.Println(p.cell)

	p2 := Person{
		firstName: "Jane",
		lastName:  "Doe",
	}
	p2.age = 27
	p2.address.city = "Bandung"
	p2.address.country = "Indonesia"

	fmt.Println(p2.address.city, p2.address.country)

	// anonymous struct
	user := struct {
		username string
		email    string
	}{
		username: "JohnDoe",
		email:    "johndoe123@gmail.com",
	}
	fmt.Println(user.email)
}

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	HomePhoneCell
}

func (person Person) fullName() string {
	return person.firstName + " " + person.lastName
}

func (person *Person) incrementAgeByOne() {
	person.age++
}

type Address struct {
	city    string
	country string
}

type HomePhoneCell struct {
	home string
	cell string
}
