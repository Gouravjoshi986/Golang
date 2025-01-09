package main
import "fmt"
import "unicode/utf8"

func dataTypes(){
	var intNum int8 = 21
	fmt.Println(intNum)

	var floatNum float32 = 10.123
	fmt.Println(floatNum)

	var result float32 = floatNum + float32(intNum)

	fmt.Println(result)

	var a int = 3
	var b int = 2
	fmt.Println(a/b) // -- 1 output 
	fmt.Println(a%b) // -- 1 output 

	var myString string = "Hello \n World"
	fmt.Println(myString)
	fmt.Println(utf8.RuneCountInString("hello"))

	var myRune rune = 'a'
	fmt.Println(myRune) // this is weird 

	var myBool bool = false
	fmt.Println(myBool)

	var myVar = "text" 
	fmt.Println(myVar)

	myVar2 := "type"       // do not do this with function . bad practice (cant understand return type of func)
	fmt.Println(myVar2)

	var3,var4 := 2,4
	fmt.Println(var3,var4)

	// We cant change a value after it is declared
	const myConst string = "constant"
	fmt.Println(myConst) // it cant be just defined ,you must initialize it also 
}
