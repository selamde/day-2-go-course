package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Welcome")
	//if statment
	age :=25
  if age > 18{
	fmt.Println("You are allowed to vote")
  }else if age > 10 {
   fmt.Println("You are a teenager")
  }else{
	fmt.Println("You are too young")
  }


  if price:=100;price>200{
	fmt.Println("This product is very expensive")
  } 
 myNemAge,err := calAge(2026)

 if(err != nil){
	fmt.Println("Ecounterd Error " , err)
	return
 }
 fmt.Println(myNemAge)


}


func calAge(birthYear int)(int, error){
  //2026 -2026
  if birthYear == 2026 {
	return 0, errors.New("Birth Year can not be the same as current year")
  } 

  return 2026 - birthYear, nil

}