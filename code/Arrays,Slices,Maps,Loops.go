package main 
import "fmt"

func array(){
	var arr [3]int32 
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr[0])
	fmt.Println(arr[1:3]) 

	// we can print address of elements also 
	fmt.Println(&arr[0])

	var arr2 [3]int32 = [3]int32{4,5,6}
	fmt.Println(arr2)

	arr3 := [3]int32{7,8,9}
	fmt.Println(arr3)
}

func slices(){
	var intSlc []int32 = []int32{1,2,3}
	fmt.Println(intSlc)
	
	fmt.Printf("the length is %v with capacity %v",len(intSlc),cap(intSlc))
	
	intSlc = append(intSlc,4)

	fmt.Printf("the length is %v with capacity %v",len(intSlc),cap(intSlc))

	// we can add one slice into another using the spread operator 
	intSlc2 := []int32{5,6}
	intSlc = append(intSlc,intSlc2...)
	fmt.Println(intSlc)

	// we can also make a slice using make function specifying data type ,len ,cap
	var intSlice3 []int32 = make([]int32,3,5)
	fmt.Println(intSlice3)

}

func maps(){
	var myMap map[string]int8 = make(map[string]int8)
	fmt.Println(myMap)

	var myMap2 = map[string]int8{"Adam":20,"Sarah":18}
	fmt.Printf("The age of Adam is %v",myMap2["Adam"])
	
	myMap3 := map[string]int8{"Sarah":18}
	fmt.Println(myMap3["Sarah"])

	fmt.Println(myMap2["Jason"]) // --returns 0 
	// to work around this we can use 
	var age1,ok1 = myMap2["Adam"]
	var age2,ok2 = myMap2["Jason"]

	if ok1{
		fmt.Printf("Adam's age is %v",age1)
	} else {
		fmt.Println("Invalid name")
	}

	if ok2{
		fmt.Printf("Jason's age is %v",age2)
	} else {
		fmt.Println("Invalid name")
	}

	// map returns a boolean value along with real value to specify if the given key is in the map or not

	delete(myMap2,"Adam") 

}

func Loops(){

	myMap := map[string]int8{"1":1,"2":2}

	myMap["1"] = 3
	myMap["2"] = 4

	myArr := [3]int8{1,2,3}
	myArr[0]=4

	mySlc := []int8{4,5,6}
	mySlc[0] = 7

	for name:= range myMap{
		fmt.Printf("Name: %v\n",name)
	}

	for i,v:= range myArr{
		fmt.Printf("Index: %v Value: %v\n",i,v)
	}

	for i,v:= range mySlc{
		fmt.Printf("Index: %v Value: %v\n",i,v)
	}
	
	var i int = 0 
	// we can use while loops like this 
	for i>10{
		fmt.Println(i)
		i++
	}
	// or 
	for{
		if i>10{
			break
		}
		fmt.Println(i)
		i++
	}

	// the same thing can be achieved using this syntax 
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
} 