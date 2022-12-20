package main

import "fmt"

func main() {
	// // declare variable
	// var greetingText string

	// // assign variable
	// greetingText = "Hello World"

	// out put
	// fmt.Println(greetingText)

	// static variable
	// var userName string = "Alireza Fazeli"
	// userName = "James Clear"

	// fmt.Println(userName)

	// // dynamic variable
	// userNameTwo := 22
	// userNameTwo = 8

	// fmt.Println(userNameTwo)

	// var userAge int = 18

	// userAgeTwo := 20

	// fmt.Println(userAge)
	// fmt.Println(userAgeTwo)

	// adminAge := userAge + 10
	// fmt.Println(adminAge)

	// adminAge = 10 + 30

	// fmt.Print(adminAge * 8)

	// var userName string =

	// var userName string = "Alireza Fazeli"
	// userName := "Alireza Fazeli"

	// var userAge int = 18
	// userAge := 18
	// // userAge = 12.2 // cant
	// // userAge = 12 // can
	// // userAge = "sda" // cant
	// fmt.Println(userAge)

	// var score float64 = 12.33222

	// fmt.Println(score)

	// var numberOne int = 18
	// var numberTwo float64 = 2.5

	// var score float64 = float64(numberOne) / numberTwo // data should be convert with same variable type

	// fmt.Println(score)

	// userName := "Hi"
	// userName = 12 // i cant do that like js or python

	// // or

	// userAge := 18
	// userAge = 18.5 // i cant do that like js or python

	// fmt.Println(userName)
	// fmt.Println(userAge)

	// var runeValue rune = 'ص'
	// var byteValue byte = '9'

	// // fmt.Println(string(runeValue))
	// // fmt.Println(string(byteValue))
	// fmt.Println(runeValue)
	// fmt.Println(byteValue)

	// var userName string = "Alireza"
	// userName := "Alireza"

	// var runeValue rune = 'ا'
	// var byteValue byte = 'g'

	// fmt.Println(string(runeValue))
	// fmt.Println(string(byteValue))

	firstName := "Alireza"
	lastName := "Fazeli"
	userAge := 18
	// fmt.Println(firstName + " " + lastName)

	// newNumber := int(2.1)

	// fmt.Println(newNumber)

	// printf -> string formatting

	// fmt.Printf("Hello my name is %v %v and i have %v years old", firstName, lastName, userAge)
	// fmt.Printf("\n Type of firstName : %T and Type of Age : %T", firstName, userAge)

	// fullName := fmt.Sprintf("Hello Guys - My Names %v %v and i have %v years old", firstName, lastName, userAge)
	fullName := fmt.Sprintf("%v %v %v", firstName, lastName, userAge)

	fmt.Printf(fullName)

}
