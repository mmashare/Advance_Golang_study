package practice

// generic strudy in golang
// generic give three type of benefits to us
// 1-> type parameter with constraint   // it's the simple generic func
// 2-> type inference                   // (~) help to find base datatype of datatype made by (type <faketypename> int) -> int is base type hrer
// 3-> type set                         // make a generic function which support many data type

import "fmt"

/////// non-generic function

func Hola(a string, b string) string {
	return a
}

// yaha hamne simple function me ek array add kiya jiske ander apne datatype ko alag se ek anem diya hai(h) and usse ek data type diya hai ,
//   and because uss h ka base type string to abb ham usse har uss parameter ke datatype se replace kar sakte hai

//////  simple generic function

func Hola2[h string](a h, b h) h {
	return a
}

///// type inference (support base type)

type Hhh string /// yaha hamne ek alag se datatype banaya hai apna but although iska base type string hi hai.

func Hola3[h ~string](a h, b h) h { // if you don't add (~) here than it do't detect base type of hhh datatpe and give us a error ,
	return a // after add (~) to string now it detect base type if you used this hhh in any variables that's are the prarmeter of this func
}

//// in go (interface{} === any) they both are same and they both are accept all datatype we can use also any, but use interface and any as a datatype in function is not good
////    because maybe they can prevents error at some time.
////   but if we want to many data type in this func then how can we give that.

// func hola4[h interface{}](a, h, b h) h {
// 	return a
// }

/// type set   (give many datatype in generic fun)

type Jiller interface {
	int | float64 /// include as many datatype as we want.
}

func Hola5[h Jiller](a h, b h) h {
	return a + b
}

func Practice4() {

	d := Hola("hii", "himasnhu") // non-generic func
	fmt.Println(d)

	r := Hola2[string]("yoo", "rakhi") // <funName>[datatype](parameter) it's called type parameter with constraint.
	fmt.Println(r)

	//    or

	r1 := Hola2("killer", "inshul") // ye bhi simple generic func hi hai, ye bhi same work karega because yaha ye auto datatype le lega.
	fmt.Println(r1)

	var t Hhh = "ki" // because we use (~) in func type now it detect hhh type base type and if the base type match with func type then ie allow to add in it.

	var t1 Hhh = "kiya" // because we use (~) in func type now it detect hhh type base type and if the base type match with func type then ie allow to add in it.

	t2 := Hola3(t, t1) // we don't need to give <funcName>[]() ----> <funcName>()
	fmt.Println(t2)

	var t3 = 3.5
	var t4 = 3.5
	var jin = Hola5(t3, t4)           // always know both have int or both varibale have float64 do'nt give eaither of them different datatype because then u need to use type casting becasue (float64 int me add nhi hota hai with type changing)
	fmt.Println(jin)

}
