package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Interpolate interface using reflect")

	var i interface{}
	i = 1

	// interface has 2 hidden field
	// type
	// value
	t := reflect.TypeOf(i)  // returns the value of type field .. equivalent to %T
	v := reflect.ValueOf(i) // returns the value of value field .. equivalent to %v
	fmt.Printf("type: %v, value: %v\n", t, v)
	fmt.Println()

	//Print type and value for different type using reflect pkg
	//PrintAllType()

	// Print field name and type of struct fields
	//PrintStruct(Student{"John", 21})

	// How to handle struct pointer?
	PrintStructPointer(&Student{"John", 21})

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

//using reflect
func Println_1(i interface{}) {
	t := reflect.TypeOf(i)  // returns the value of type field .. equivalent to %T
	v := reflect.ValueOf(i) // returns the value of value field .. equivalent to %v or i.(type)
	fmt.Printf("type: %v, value: %v\n", t, v)
}

func PrintStruct(s interface{}) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct {
		fmt.Print("Not of type struct")
		return
	}
	v := reflect.ValueOf(s)
	numberOfFields := t.NumField()
	for i := 0; i < numberOfFields; i++ {
		fName := t.Field(i).Name
		fType := t.Field(i).Type
		fValue := v.Field(i)
		fmt.Printf("Field %v, Name: %v, Type: %v, Value: %v\n", i, fName, fType, fValue)
	}

}


func PrintStructPointer(s interface{}) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Ptr {
		fmt.Print("Not of type pointer")
		return
	}

	// If its of pointer type then we need to get the value of that pointer which is the struct in our case
	v := reflect.ValueOf(s)
	fmt.Printf("v: %v\n", v)

	rv := reflect.Indirect(v)
	fmt.Printf("rv: %v\n", rv)
	fmt.Print(rv.FieldByName("Nam"))
	fmt.Print("\n")

	//PrintStruct(rv) This won't work as desire because we need to pass imterface value to our function not reflect.value
	// Lets convert reflect.value to interface type
	PrintStruct(rv.Interface())
}