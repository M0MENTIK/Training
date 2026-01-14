package main

import "fmt"

func main() {

	fmt.Println()
	fmt.Println("-------------TASK 1---------------")
	//Task 1
	max := Dog{}
	AllSpeak(max)
	bob := Cat{}
	bob.Speak()
	var fart Speaker
	fart = Dog{}
	fart.Speak()
	fmt.Println()
	//

	fmt.Println()
	fmt.Println("-----------------TASK 2-------------------")
	//Task 2
	var shapes = []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 3},
		Rectangle{Width: 5, Height: 3},
		Rectangle{Height: 18, Width: 21},
		Circle{Radius: 12},
		Circle{Radius: 13},
	}
	var SumArea float64
	for _, a := range shapes {
		SumArea += a.Area()
	}
	fmt.Println(SumArea)
	fmt.Println()

	//

	fmt.Println()
	fmt.Println("-----------------TASK 3-------------------")
	//Task 3
	var items = []Printer{
		Document{DocumentText: "word"},
		Image{ImageText: "image"},
	}
	PrintAll(items)
	fmt.Println()

	//

	fmt.Println()
	fmt.Println("-----------------TASK 4-------------------")
	//Task 4
	var card1 = CreditCard{
		cardNumber: "4354 3453 3453 3463",
	}
	card2 := CreditCard{
		cardNumber: "1322 2342 2343 3244",
	}

	var paypal1 = PayPal{
		email: "paypal1@gmail.com",
	}
	paypal2 := PayPal{
		email: "paypal2@gmail.com",
	}
	ProcessPayment(card1, 12)
	ProcessPayment(paypal1, 45)
	ProcessPayment(card1, 456)
	ProcessPayment(card2, 56)
	ProcessPayment(paypal2, 3333.345)
	fmt.Println()

	//

	fmt.Println()
	fmt.Println("-----------------TASK 5-------------------")
	//Task 5
	DescribeSwitch(4)
	DescribeSwitch("123")
	DescribeSwitch(Person{Name: "Alex"})
	DescribeSwitch(3.4)

	DescribeAssertion(4)
	DescribeAssertion("123")
	DescribeAssertion(Person{Name: "Alex"})
	DescribeAssertion(3.4)
	fmt.Println()

	//

	fmt.Println()
	fmt.Println("-----------------TASK 6-------------------")
	//Task 6
	Process(&File{}, "TestText")
	fmt.Println()
}
