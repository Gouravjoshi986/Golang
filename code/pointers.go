package main 
import "fmt"
// pointers are special variable which stores address . 
// defined by an asterisk * 
func pointers(){
	var p*int32 = new(int32)
	var i int32
	// we can use method new(int32) to give the pointer a new address to point to with default value of int32 

	// to derefrence the pointer we use *p 
	*p = 10 // to change the value of location that p is pointing to 
	p = &i
 	*p = 5 

	// slices also show this behaviour 
	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)    		// 1,2,4
	fmt.Println(sliceCopy)  	// 1,2,4
	// changing copy of slice also changed slice 

}

func square(thing2 [3]float64)[3]float64{
	for i:= range thing2{
		thing2[i]=thing2[i]*thing2[i]
	}
	return thing2
}

var thing1 = [3]float64{1,2,3}
var result [3]float64 = square(thing1)

// here The thing is passed using pass by value : we are using double memory 

// to tackle this we can use pointers 
// : square(&thing1)  and square(thing2 *[3]float64)  would solve this. 
