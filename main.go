/// type conversion advance study

package main

import (
	practice "advancegoconcept/practices"

	"fmt"
)

// type conversion with interfaces in go
// type conversion means change variable data type to other data types. we already study it but here we study type conversion in interfaces.

type myint int   	// we create our datatype name of (myint) but the reality is it made on int data-type means it's base data-type is (int).

func (i myint) myintMethod() myint {             ////  method of myint type
	return i + 1  								//// jo bhi value yaha ayegi vo +1 ho jayegi
}



func main(){

	var g myint = 34   // yaha hamne iss variable ko myint name ka datatype diay hai.
	fmt.Println(g.myintMethod())

	var u = int(g)  // yaha hamne ek myint ko ki ek fake datatype type tha jiska base datatype int tha usse int me convert kar diya. 
	fmt.Printf("%T", u)

	practice.Practice()
	practice.Practice2()
	practice.Practice3()
	practice.Practice4()
	practice.Practice5()
	practice.PointerStudy()
}