package internal

import "fmt"

type Customer struct {
	Id string
}

func (c *Customer) Update(itemName string) {
	fmt.Printf("Enviando email para " + c.Id + " sobre o item " + itemName + " estar em estoque. \n")
}

func (c *Customer) GetId() string {
	return c.Id
}
