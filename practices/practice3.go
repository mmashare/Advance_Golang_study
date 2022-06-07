package practice

/// study of reflect pkg in  go.
/// it very complex thing in go but at the same time is easy.
/// it mainly use to find type of var at run time in go. normally we find type at compile time in go
/// feature of this pkg.
/// 1 -> give concrete type and baseType
/// 2 -> give value of interface adn struct [basically vaiable]
/// 3 -> give Numfeild {hoe many keys in struct} ||| warning -> this method only work with struct.

import (
	"fmt"
	"reflect"
)

type Order struct {
	FighterId   int64
	FighterName string
}

func Killer(u Order) string { // here we use (Order) type because it's the real type of that struct
	g := fmt.Sprintf("FighterId -> %T \n FighterName -> %T", u.FighterId, u.FighterName) // yaha ham compile time par type checking kar rahe hai
	return g             // here we find type of vailble but in compile time.
}

func Killer2(l interface{}) {
	a := reflect.TypeOf(l) // it return a concrete type of var (Order is concrete and struct is base type)
	fmt.Println(a)

	b := reflect.ValueOf(l) // it return values of struct(Order)
	fmt.Println(b)

	c := b.Kind()  // it return the real variable type(or u can say Base Type of var) which is {{{{struct}}}
	fmt.Println(c)

	v := b.NumField()   // it return how many (key) in the struct which is also called NumFeild in reflect pkg. warning -> {NumFeild() and Feild() method only works with stuct in reflect pgk.}
	fmt.Println(v)       // hamne yaha b isliye liya hai because hamme sarii ki count chiye aur vo b ke pas hai because uske pas sari key ki value hai. that's why we select b.

	

}

func Practice3() {
	y := Order{FighterId: 23, FighterName: "khabib"}

	k := Killer(y)
	fmt.Println(k)

	Killer2(y)

}