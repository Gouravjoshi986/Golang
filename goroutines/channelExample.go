package main 
import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_VEG_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func home(){
	var vegChannel = make(chan string)
	var tofuChannel = make(chan string)

	var websites = []string{"walmart.com","bigbasket.com","swiggy.com"}

	for i:= range websites{
		go fetchVegPrices(websites[i],vegChannel)
		go fetchTofuPrices(websites[i],tofuChannel)
	}
	sendMessage(vegChannel,tofuChannel)
}

func fetchVegPrices(website string,c chan string){
	for{
		time.Sleep(time.Second*1)
		var vegPrice = rand.Float32()*20
		if vegPrice < MAX_VEG_PRICE{
			c <- website
			break
		}
	}
}

func fetchTofuPrices(website string,c chan string){
	for{
		time.Sleep(time.Second*1)
		var tofuPrice = rand.Float32()*20
		if tofuPrice < MAX_TOFU_PRICE{
			c <- website
			break
		}
	}
}

func sendMessage(vegChannel chan string,tofuChannel chan string){
	select{
	case website := <- vegChannel:
		fmt.Printf("\n Found veg deal at %v",website)
	case website := <- tofuChannel:
		fmt.Printf("\n Found Todu deal at %v",website)
	}
}

// this select statement will look for a result from any of the channels . if it gets one it will execute the returned case and exits. 