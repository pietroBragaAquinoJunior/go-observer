package internal

import "fmt"

type Item struct {
	name         string
	inStock      bool
	observerList []Observer
}

func NewItem(name string) *Item {
	return &Item{name: name}
}

func (i *Item) UpdateAvaliability() {
	fmt.Printf("O item está em estoque agora. \n")
	i.inStock = true
	i.NotifyAll()
}

func (i *Item) Register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) DeRegister(o Observer) {
	i.observerList = RemoveFromList(i.observerList, o)
}

func (i *Item) NotifyAll() {
	for _, observer := range i.observerList {
		observer.Update(i.name)
	}
}

func RemoveFromList(list []Observer, element Observer) []Observer {
	listLen := len(list)
	for i, observer := range list {
		if observer.GetId() == element.GetId() {
			list[listLen-1], list[i] = list[i], list[listLen-1]
			return list[:listLen-1]
		}
	}
	return list
}
