package practice

import (
	"fmt"
)

// here we study about pointer in go
// & -> this provide pointer(location of that var in hash type)
// * -> this used after you put a (&)pointer to a var (if cannot used without & means it cannot use before &pointer you have to use it after you put a pointer to a var)
// & - fmt.Println(&) it's only print location of a var in hex
// * -> fmt.Println(*) it's provide the real value of pointer var that u provide in a var

func PointerStudy(){

	var hms int = 9

	kk2 := hms
	fmt.Println(kk2)        // it's print 9 but it's not the orignal 9(value) , it just a reference of real value 9

	kk := &hms        // here u give a pointer of hms var (store location in hex)
	fmt.Println(kk)   // (0xc00000e040)it only print the location of hms in hex not the value
	fmt.Println(&kk)  // (0xc00000e040) it also print location of hms not the value of hms
	fmt.Println(*kk)  // (9) here it print the value of hms and it's real , so after u give a pointer of var then u have to use * to print the value of var


	
	var f int = 90
    var p *int     // declare the pointer

	// here u give a pointer of int it's mean , u say that i will only store the pointer of var in that not the simple value of int(7,67,45,23) becasue it's not work here
	// it's show error if u do that though u  do this *int it's mean u say i will only store which type is int and also i can only store &pointer of that var
	// not the normal refrence 

    p =   &f       // initialize the pointer
	// here u store &(pointer) of int type var
    fmt.Println(*p)  // 90
	// *p -> it print value of f here



	/////// giving pointer to pointer ////////////

	var hi int = 90

	b := &hi
	c :=&b
	fmt.Println("-----------------------------")
	fmt.Println(b) // ->> 0xc0000140c8 (as always) yaha jo address hai vo hi var ka hai {yaha ye jo address hai vo permanent nhi b ka because ka kud ka alag address hai ussne as a value (hi)var ka address store kiya hai}
	fmt.Println(c) // ->> 0xc00000e018 (as always)  yaha jo address hai vo real add hai b var ka

	fmt.Println(*b)      // --- >> 90 

	fmt.Println(*c)  // -->> 0xc0000140c8 (why) because yaha ham ek pointer se pointer liya hai iss var me that's why yaha vo only location de paya uss var ki


		///////////////// pointer to interface ///////////////////

	var df interface{
		
	}

	fff := &df
	
	fmt.Println(*fff)  /// -> nil (because interface to empty hai that's why)                                                                                                
	fmt.Println(fff)   // -> 0xc0000105b0 (as always it give location in hex)



	/////////////////// pointer use with func ///////////////////

	bn := "hms "  
	Myname(&bn)  // -> hms (becasue func ke ander *pointer se value print ho rahii hai)

	/////////////// func which return a pointer of string //////////////////////////

	jin:= Mysecondname()
	fmt.Println(jin)     // -> 0xc0000105d0 (yaha func ke ander kke var ka memory location hai)
	fmt.Println(*jin)    // -> killer ground (yaha uss var ki value dega return me)


	////// always know that h:= (&hj + 78) app asa nhi kar sakte hai go lang me app pointer ko matchmatically sign calculation nhi kar sakte hai.

}

func Myname(a *string){  // (a *string) now u declar that u only takes a var which can give it's pointer in func() as an argument, no normall argumetment will work.  
	fmt.Println(*a)
}

func Mysecondname() *string {
	th := "killer ground"
	
	return &th
}