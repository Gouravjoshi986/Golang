package main
import "fmt"
// When we are not sure or we want to write a single piece of code for multiple data types we can use generics. 

// we also have "T any" - for any type but using this comes with conditions as not every data type is complatible with every operation (eg booleans are not compatible with arithmetic)

// We can also use generics with structs - generally when we load data from APIs

func gen(){
	var intSlc = []int{1,2,3}
	fmt.Println(sumSlc[int](intSlc))

	var fltSlc = []float32{1,2,3}
	fmt.Println(sumSlc[float32](fltSlc))
}

func sumSlc[T int | float32 | float64](slice []T) T{
	var sum T
	for _,v := range slice{
		sum+=v
	}
	return sum
}