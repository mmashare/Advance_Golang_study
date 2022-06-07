package practice

// study of type switch . it's an advance version of type asseration.

import "fmt"

func Practice2() {
	var x interface{} = true

	switch v := x.(type) {
	case nil:
		fmt.Println("\nx is nil") // here v has type interface{}
	case int:
		fmt.Println("\nx is", v) // here v has type int
	case bool:
		fmt.Println("\nx is bool ") // here v has type interface{}
	case string:
		fmt.Println("\nx is string")
	default:
		fmt.Println("\ntype unknown") // here v has type interface{}
	}

}
