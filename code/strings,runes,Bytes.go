package main 
import (
	"fmt"
	"strings"
)
func stringFunc(){
	var str string = "resume"
	// we can access this using an index but the problem is, we cant print it. 
	fmt.Println(str)
	var indexed = str[0]
	fmt.Printf("%v %T",indexed,indexed) //returns 113,uint8 
	for i,v := range str{
		fmt.Printf("Index: %v Value: %v",i,v)
	} // Here an array of Bytes is Printed 
	// If a character takes 2 or more bytes then index could be skipped . 

	// this is as go uses utf 8 encoding to represent strings 
	// solution - use array of runes 
	var str2 = []rune("resume")
	for i,v := range str2{
		fmt.Printf("Index: %v Value: %v",i,v)
	}
	// here the unicode number value of the bytes is printed so we dont get skipped indices. 
	// ie values of those bytes are printed
	 
	// runes can be declared using single quotes 
	var myRune = 'a' // prints 97 
	fmt.Println(myRune)
	 
	// strings are immutable in go. cant modify after they are created 
	var strSlice = []string{"c","a","t","f","i","s","h"}
	var strConcat = ""
	var strBuilder strings.Builder
	for i:= range strSlice{
		strConcat+=strSlice[i]
		strBuilder.WriteString(strSlice[i])
	}
	
	var strNew = strBuilder.String()
	fmt.Println(strConcat) // catfish
	fmt.Println(strNew) // catfish

	// This is inefficient so we use go's string package for string building 

}