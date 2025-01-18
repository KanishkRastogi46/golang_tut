package main

import (
	"fmt"
	"time"
)

type customer struct {
	id int
	name string
	email string
}

type order struct {
	id int	
	name string
	status string
	orderedAt time.Time
	customer // embedding customer struct
}

// constructor function
func newOrder(id int, name string, status string) *order{
	// initialize the struct
	nOrd := order{
		id: id,
		name: name,
		status: status,
		orderedAt: time.Now(),
	}
	return &nOrd
}

// receiver function
func (o *order) changeStatus(status string) string {
	o.status = status
	return o.status
}

func main() {
	var ord1 order = order{
		id: 1,
		name: "Kanishk",
		status: "received",
		orderedAt: time.Now(),
	}

	fmt.Println("order 1: ")
	fmt.Println(ord1.id)
	fmt.Println(ord1.name)
	fmt.Println(ord1.status)
	fmt.Println(ord1.orderedAt)
	ord1.changeStatus("shipped")
	fmt.Println(ord1)

	fmt.Println("order 2: ")
	ord2 := newOrder(2, "Krishna", "delivered")
	fmt.Println(ord2.id)
	fmt.Println(ord2.name)
	fmt.Println(ord2.status)
	fmt.Println(ord2.orderedAt)
	fmt.Println(*ord2)

	// if struct is only used once, we can use anonymous struct
	language := struct{
		name string
		version float64
	} {"Golang", 1.23}
	fmt.Println(language)

	newCustomer := customer{
		id: 1, 
		name: "Paul",
		email: "paul@gmail.com",
	}
	newOrd := order{
		id: 3,
		name: "Paul",
		status: "received",
		orderedAt: time.Now(),
		customer: newCustomer,
	}
	fmt.Println(newOrd)

	orders := []order{
		{1, "Kanishk", "received", time.Now(), customer{1, "Kanishk", ""}},
		{2, "Krishna", "shipped", time.Now(), customer{2, "Krishna", ""}},
	}
	fmt.Println(orders)

	var myOrder *order = newOrder(3, "Paul", "delivered")
	orders = append(orders, *myOrder)
	fmt.Println(orders)
}