package main

import "fmt"

type Interface interface {
}

func DescribeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("This is number:", v)
	case string:
		fmt.Println("This is string:", v)
	case Person:
		fmt.Println("This people:", v.Name)
	default:
		fmt.Println("Unknown type")
	}
}

func DescribeAssertion(i interface{}) {
	if v, ok := i.(int); ok {
		fmt.Println("This is number:", v)
	} else if v, ok := i.(string); ok {
		fmt.Println("This is string:", v)
	} else if v, ok := i.(Person); ok {
		fmt.Println("This people:", v.Name)
	} else {
		fmt.Println("Unknown type")
	}
}

type Person struct {
	Name string
}
