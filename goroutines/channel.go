// channel is a way to enable a goroutine to pass around information.package goroutines

// They hold data 
// they are thread safe (avoid races)
// we can listen when data is added or removed in a channel and we can block execution until one of these executions happen

package main 
import (
	"fmt"
	"time"
)
// you can make it using make(chan dataType)
// we use <- to add values . 
// you can think of it having an underlying array 
func channel(){
	var c = make(chan int)
	c <- 1
	var i = <- c // this is how you retrieve value 
	// After retreving a value from the channel. the value gets popped out.and channel gets empty. 
	fmt.Println(i)
}
// runnning this will give us a deadlock error as runnning c<-1 will lock the thread and it will not run past it. So go throws an error 

// Channels are designed to work with goroutines
func work(){
	var c = make(chan int)
	go process(c)
	for i:= range c{
		fmt.Println(i)
	}
}
func process(c chan int){
	// c <- 123 
	defer close(c)
	for i:=0;i<5;i++{
		c <- i
	}
	// close(c)
}
// If we run it without close(c) "closing the channel" it will again give us a deadlock error
// Because after all iterations it will again look for any element in channel and lock the threads
// we can also use the defer statement : which means that defer this task till the function ends

// What we did till now is called an unbuffered channel and it can only store one value. 
// So due to this the process function will not exit unless the main function completes all of its work and channels close - we can optimise this using buffered channel. 

// Now lets see Buffered Channel 

func bufferedChan(){
	var c = make(chan int,5)  // buffered channel
	go process2(c)
	for i:= range c{
		fmt.Println(i)
		time.Sleep(time.Second+1)
	}
}
func process2(c chan int){
	defer close(c)
	for i:=0;i<5;i++{
		c <- i
	}
	fmt.Println("Exiting process")
}

// now the process2 function will exit after completing its task and adding elements to channel
// the buffered chan func will continue doing its work
