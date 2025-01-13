package main 
import "fmt"

type animal struct{     // this is how you 
	name string            // create structs 
	speed uint8
	Info dog    // we can use this  OR 
	dog         // we can use this field directly 
}                // this way we can skip writing info while accessing . ie - dog1.breed

func (e animal) greetName()string{
	var greet = "hello"+e.name
	return greet
}
// this is a method that we created for animal struct : (e animal) this assigns it to animal structure. so it could access all the props

type dog struct{
	breed string
	harmful bool
}
func Structs(){
	var dog1 = animal{"jerry",6,dog{"indian",false},dog{"indian",false}}

	fmt.Println(dog1.name,dog1.speed,dog1.Info.breed,dog1.Info.harmful,dog1.breed,dog1.harmful)
	
	// we can do this with any time . we can write int -- now this is a variable called int with type also int 
	
	// we can also define structs anonymously but we have to initialise it then and there only

	var myAnimal = struct{
		name string
		age uint8
		breed string
	}{"tom",8,"cat"}

	fmt.Println(myAnimal.name,myAnimal.age,myAnimal.breed)
	// But this is NOT REUSABLE 

	fmt.Println(dog1.greetName())	
}

// now coming to interface 
type gasEngine struct{
	mpg uint8
	gallons uint8
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

func (e gasEngine) milesLeft()uint8{
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft()uint8{
	return e.kwh*e.mpkwh
}

type engine interface{
	milesLeft() uint8 
}

// Now that we are using interface as a params we can take both engines as a parameter as long as it has a method called milesLeft() which returns a uint8 quantity 
func canMakeIt(e engine,miles uint8){
	if(miles<=e.milesLeft()){
		fmt.Println("you can make it")
	}else{
		fmt.Println("you cant make it")
	}
}