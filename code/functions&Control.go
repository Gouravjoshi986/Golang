package main
import (
	"fmt"
	"errors"
)

func printMe(param string){
	fmt.Println("Hello "+param)
}

func intDivision(numerator int,denominator int) (int,int,error){

	var err error //error is also a data type
	if denominator==0{
		err = errors.New("Cannot Divide by zero")
		return 0,0,err
	}

	if 1==1 && 2==2{  // && || etc 
		fmt.Println("Passed two checks")
	}

	var result int = (numerator/denominator)
	var remainder int = (numerator%denominator)
	return result,remainder,err
}



