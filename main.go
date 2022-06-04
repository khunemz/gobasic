package main

import (
	"fmt"
	"gobasic/customer"
	"gobasic/user"
	"unicode/utf8"
)

func main() {

	person := user.Person{FirstName: "Chutipong", LastName: "Roobklom", Age: 32}
	fmt.Printf("Name: %v %v,Age: %v\n", person.FirstName, person.LastName, person.Age)

	fmt.Printf("GET Name: %v \n", person.GetName())
	person.SetName("chutipong", "roobklom")
	fmt.Printf("GET Name Again: %v \n", person.GetName())

	var num1 int
	var num2 *int
	num1 = 10
	num2 = &num1

	fmt.Printf("%v \n", &num1)
	fmt.Printf("%v \n", num2)
	fmt.Printf("%v \n", *num2)

	res := customer.Sum(&num1)
	fmt.Printf("result is : %v\n", res)
	// fmt.Printf("%v", summaryvar)

	println("Hello world")
	fmt.Printf("Hello %v\n", "Hello world another one.")

	fmt.Printf("Customer : %v\n", customer.GetCustomer())

	myNumber := 1

	fmt.Printf("Number: %v\n", myNumber)

	point := 72

	grade := ""
	if point < 50 && point >= 0 {
		grade = "F"
	} else if point >= 50 && point < 60 {
		grade = "D"
	} else if point >= 60 && point < 70 {
		grade = "C"
	} else if point >= 70 && point < 80 {
		grade = "B"
	} else if point >= 80 {
		grade = "A"
	} else {
		grade = "INVALID"
	}

	fmt.Printf("Your grade is : %v\n", grade)

	// ARRAY
	myArr := []int{10, 20, 30, 50, 60, 70, 80, 90, 100}
	myArr = append(myArr, 100)
	fmt.Printf("MyArray: %v\n", myArr)

	for i := 0; i < len(myArr); i++ {
		fmt.Printf("Nums in arr : %v \n", myArr[i])
	}

	name := "ชุติพงศ์"
	fmt.Printf("Chars in name count: %v\n", utf8.RuneCountInString(name))

	// slice
	// from index, to index+1
	y := myArr[1:5]
	fmt.Printf("Slice : %v\n", y)

	// map
	countries := map[string]string{}
	countries["th"] = "Thailand"
	countries["usa"] = "United States"
	countries["jp"] = "Japan"

	country, isExisted := countries["jp"]
	if isExisted {
		fmt.Printf("Country existed %v\n", country)
	} else {
		fmt.Printf("Country not existed")
	}

	cnt := 0
	for cnt < len(myArr) {
		fmt.Printf("Value in Array is : %v\n", myArr[cnt])
		cnt++
	}

	// for length
	for _, v := range myArr {
		fmt.Printf("value in array is : %v\n", v)
	}

	c := sumNumber(10, 30)

	fmt.Printf("Summary number : %v\n", c)

	one, _ := tupleSum(40, 50)

	// fmt.Printf("value : %v, keyword: %v", one, two)
	fmt.Printf("value : %v", one)

	summation := cal(sumNumber)
	fmt.Printf("Summation: %v\n", summation)
}

func sumNumber(a, b int) int {
	return a + b
}

func cal(f func(int, int) int) int {
	summation := f(10, 50)
	fmt.Printf("%v\n", summation)
	return summation
}

func tupleSum(a, b int) (int, string) {
	return a + b, "Summary is :"
}
