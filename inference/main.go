package main

import "fmt"

func main(){
	var firstName, lastName, age, salary = "John", "Maxwell", 28, 50000.0

    fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T\n", 
        firstName, lastName, age, salary)
}