package main

import (
	"fmt"
	"structs/models"
)

func newPerson(p models.Person) *models.Person {
	return &p
}

func main() {
	p := models.Person{
		Name:     "Carlos",
		Lastname: "Sandoval",
		Age:      23,
	}
	node1 := models.Node[int]{Value: 1}
	node2 := models.Node[int]{Value: 2}
	node3 := models.Node[int]{Value: 3}

	// Enlazar los nodos
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = nil
	fmt.Println(p)
	fmt.Println(newPerson(models.Person{Name: "Pito Perez", Lastname: "Peraza", Age: 44}))

	tmp := &node1
	for tmp != nil {
		fmt.Println(tmp.Value)
		tmp = tmp.Next
	}
}
