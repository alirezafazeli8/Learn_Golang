package main

import "fmt"

func main() {

	// static type
	var firstName string = "Alireza"
	var lastName string = "fazeli"
	var currentYear int = 2022

	// dynamic type
	adminFirstName := "James"
	adminLastName := "Clear"
	myBirthday := 2004

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(adminFirstName)
	fmt.Println(adminLastName)

	fmt.Println(currentYear)
	fmt.Println(myBirthday)

	currentYear = currentYear + 1

	var yourAge int = currentYear - myBirthday

	fmt.Println(yourAge)

}
