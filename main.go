package main

import "observer/internal"

func main() {
	item := internal.NewItem("tapete")
	customer := &internal.Customer{Id: "pipiboy@gmail.com"}
	customer2 := &internal.Customer{Id: "anacaroline@gmail.com"}

	item.Register(customer)
	item.Register(customer2)
	item.UpdateAvaliability()
}
