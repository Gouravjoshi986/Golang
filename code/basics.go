package main
import "fmt"

func main(){
	fmt.Println("hello world from Basic file")
	// dataTypes()
	printMe("Gourav")

	var result,remainder,err = intDivision(10,3)
	
	// IF Else statements 
	if err!=nil{
		fmt.Println(err.Error())
	} else if remainder==0{
		fmt.Printf("The result of integer division is %v",result)
	} else{
		fmt.Printf("The result of integer division is %v with remainder %v",result,remainder)
	}

	// Switch statements
	switch{
		case err!=nil:
			fmt.Printf(err.Error())
		case remainder==0:
			fmt.Printf("The result of integer division is %v",result)
		default:
			fmt.Printf("The result of integer division is %v with remiander %v",result,remainder)
	}

	switch remainder{
		case 0: fmt.Printf("Division was Exact")
		case 1,2 : fmt.Printf("Division was close")
		default: fmt.Printf("Division was not close") 
	}
}