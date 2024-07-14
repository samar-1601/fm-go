package main

import "fmt"

// kind of a go version to OOPs,
// uses composition over inheritance,
// Go doesn't have classes

// declaring a struct
// since it is kinda type of typescript, go compiler doesn't give us
// warnings for unused variables
type Person struct {
	Name string
	Age  int
}

func changeName(p *Person) {
	p.Name = "Aloo"
}

// tells that the changeNameAsMethod() can only be called from the Person type
func (p *Person) changeNameAsMethod(name string) {
	p.Name = name
}

func main() {
	myPerson := Person{
		Name: "Samar",
		// since Age is not assigned here, it will take the by defualt value = 0
	}

	fmt.Println(myPerson)
	// using interface to print
	fmt.Printf("before : %+v\n", myPerson) // %v signifies that it can take any values, "+" tells to mention the field as well

	changeName(&myPerson) // pass by reference

	fmt.Printf("after : %+v\n", myPerson)

	myPerson2 := Person{
		Name: "Samar2",
	}

	fmt.Printf("before : %+v\n", myPerson2)
	myPerson2.changeNameAsMethod("MethodAloo")
	fmt.Printf("before : %+v\n", myPerson2)

	// Go has similar Pointers concept as C++
}
