package practice

import "fmt"

//// we study Type assertions here .
/// basically type assertions and type switch is a process to find interface value & datatype.

func Practice() {
	fmt.Println("Hello")

	var f interface{} = "hms" // here we make interface in which we put a string value in it.

	var d = f.(string) // yaha ham .(string) se ye find kar rahe hai ki ager f interface me string ki koi value hai to vo value hamee dee.
	fmt.Println(d)

	var q, ok = f.(string)     // q = value and ok = true,false (ager value string hogi too true ayega and ager value string nhi hogi to false answer agyega)
	fmt.Printf("%v %v", q, ok) // second variable (ok) is temprary

	var b, n = f.(int)   // ye statement false hai because f interface me string ki value hai int ki nhi that's why b variable me 0 value ayegi and n variable me false value ayegi.
	fmt.Printf("\n%v %v", b, n)
}