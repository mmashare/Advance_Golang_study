package practice

import (
	"fmt"
	"io"
	"os"
)

// io package is same like fs packge in node js.
// os package hamre system me file open karne and system me empty file create karne me help karta hai.

func Practice5(){

	file, _ := os.Create("damn.txt")
	
	////// write a file by I/O package //////////

	w := io.Writer(file)
	r,err1 := w.Write([]byte("hellooooooooo"))
	
	////// or //////////

	t,err2 := io.WriteString(w," \nsecond method of write string")

	fmt.Println(t,err2)
	
	fmt.Println(r,err1)
	
	/////// read a file by I/O package /////////
	
	q,_ := os.Open("damn.txt")

	g := io.Reader(q)
	j := make([]byte,1)              // bascially is file ke ander pura content hota hai that's ham isse print karte hai and it's called buffer.
	
	

	for {
		h,err := g.Read(j)  /// yaha hamme .Read method only byte return karega file ke jo hamne make function me diye hai.
		fmt.Printf("%v %v \n%v\n",h,err,string(j))  // because h me only byte hai that's why yaha ham j(jiske ander make method hai) usse bhi print kar rahe hai jiske ander pura content hai. j variavle ko abb ham buffer bolege.

		if err != nil{         // if ager hamra code me error hai the hamra loop break ho jayega.ye err hammme .read method de sakta hai.
			break
		}

	}


	//// or ///

	/// why we make seeker here -> we make seeker here because let's take a example (manlo apne koi file ek bar coe me read karli then app koi dobara read karna chate ho to then apko second time file read karne se phale ek sekker banan padega  because if you don't make a seekerager apne koi seeker banye dobara ussi file ko read karne ki sochi to yaha ye file second me kuch bhi read nhi karegi empty read hogi because ab ye file ke last word ke badd se read karna start karega that's why)
	//// seeker ek controller ke tarah work karta hai like vo hamme batata hai ki hamme file kha se read karni hai like file ke kis work se hamme file read karani hai.

	sekker := g.(io.Seeker)   /// yaha hamne ek fiel me sekker create kara hai
	sekker.Seek(0,io.SeekStart)  // yaha hamne .ssek method me bataya hai ki hamme kis word se file read karai surur karni hai and then .SeekStartmethod suru kar diya hai

	/// type of .Seek(-5,io.SeekEnd)  -> yaha ye ab last se only five value read karega


	j,err3 := io.ReadAll(g)      // now this code work normally
	fmt.Println(string(j),err3)  // if u don't change j into string then it give u some decimal number of alfabat that we don't want that's why we change j variable into string so we can read our content 

	file.Close()
}