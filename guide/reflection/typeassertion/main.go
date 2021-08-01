package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {

	//PrintAllType()

	// This will work because this has value of type string in it
	//TypeAssert("John")

	// compiler won't throw error as func expect empty interface type
	// this will panic as it does not have string value in it
	// o/p > panic: interface conversion: interface {} is int, not string
	//TypeAssert(2)

	// How to handle above?
	// Handle gracefully using bool value, so it won't panic
	// o/p > Value is of type int not of type string
	//TypeAssertBool(2)

	// But I have so many types, its not the correct way to write multiple if-else block for each type assertion
	// Now What ??
	// TypeAssertion using Switch case
	//PrintAllTypeSwitch()

	// This also won't suffice if to be sure that we won't miss any type
	// Have to use reflect package

}

func PrintAllType() {
	Println := Println_1

	Println("string")
	Println(true)
	Println(1)
	Println(2.45)
	Println([5]int{})
	Println([]string{})
	Println(map[string]string{})
	Println(Student{"John", 21})
	Println(&Student{"John", 21})
}

// use Printf substitution
// %T for type
// %v for value
func Println_1(i interface{}) {
	fmt.Printf("Type: %T, Value:%v\n", i, i)
}

func TypeAssert(i interface{}) {
	value := i.(string) // Give me the string value from this interface type
	fmt.Printf("Value is %v\n", value)
}

func TypeAssertBool(i interface{}) {
	value, ok := i.(string) // Give me the string value from this interface type
	if ok {
		fmt.Printf("Value is %v\n", value)
	} else {
		fmt.Printf("Value is of type %T not of type string\n", i)
	}
}

func TypeSwitch(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Print(i.(string))
	case bool:
		fmt.Print(i.(bool))
	case int:
		fmt.Print(i.(int))
	case float64:
		fmt.Print(i.(float64))
	case [5]int:
		fmt.Print(i.([5]int))
	case []string:
		fmt.Print(i.([]string))
	case map[string]string:
		fmt.Print(i.(map[string]string))
	case Student:
		fmt.Print(i.(Student))
	case *Student:
		fmt.Print(i.(*Student))
	default:
		fmt.Print("Unknown type")
	}
	fmt.Print("\n")
}

func PrintAllTypeSwitch() {
	Println := TypeSwitch

	Println("string")
	Println(true)
	Println(1)
	Println(2.45)
	Println([5]int{})
	Println([]string{})
	Println(map[string]string{})
	Println(Student{"John", 21})
	Println(&Student{"John", 21})

	Println(map[int]string{})
}
