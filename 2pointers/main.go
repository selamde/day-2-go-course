package main

import (
	"fmt"
)

type Person struct{
	Name string
    Exp int
}

type Food struct{
	Name string
	Content string
	Cal float64
}

func main() {
	fmt.Println("Pointers")
	age := 20

	student := Person{
		Name:"Selam",
		Exp: 5,
}


food := Food{
	Name: "Avovcado",
	Content:"Oil substance",
	Cal: 20.67,
}

myFoodPointerVariable := &food
ModifyFoodCal(myFoodPointerVariable)


	myStudentPointerVariable := &student
  ModifyExp(myStudentPointerVariable)

    myPointerVariable := &age
	*myPointerVariable = 100
	ModifyAge(&age)

		fmt.Println(age)
		fmt.Println("Exp ", student.Exp)
		fmt.Println( "New cal ",food.Cal)

}


func ModifyAge(age *int){
	//create a new copy
	*age = 150
}


func ModifyExp(person *Person){
   person.Exp = 10

}


func ModifyFoodCal(food *Food){
   food.Cal = 35.98
}