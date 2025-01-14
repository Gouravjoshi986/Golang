// way to launch multiple functions and have them execute concurrently.
// Concurrency != Parallel execution 
// If we have one core and we schedule tasks such that when one task is in wait phase we move to another task to complete it and return back when task 1 wait is over to schedule next tasks. 

// Parallel execution is using multiple cpus to parallely execute two or more tasks. we can achieve concurrency via parallel execution.
package main 
import (
	"fmt"
	"time"
	"sync"
)


// To introduce Concurrency we use the 'go' keyword before function call. 
// But if we only used go keyword it will skip all tasks and will end before any task could be completed .. So we use wait groups. 
var wg = sync.WaitGroup{}
// wait group is just like a counter 
// Add - adds to counter 
// Wait - waits till the counter gets back to 0 
// Done - decreases the counter 

var dbData = []string{"id1","id2","id3","id4","id5"}
var results = []string{} // to store results from db

// Now this Creates Problem . Here due to concurrency there could be times where one memory location is being accessed by two processes. 
// So we shouldnt run it like this. 

// Answer - MUTEX  - fromt he sync package 
// It is very important to note that it is crucial to put the lock and unlock at the right places. 
// if we were to put it before the delay call it will completely destroy the concurrency 

// One drawback of this mutex is it completely locks all things. 
// Go provides another function called the read write mutex. with Read lock and Read unlock 

// var a = sync.Mutex{}  

var a = sync.RWMutex{}
func main(){
	t0 := time.Now()
	for i:=0;i<len(dbData);i++{
		wg.Add(1)       
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\n Total Execution time: %v",time.Since(t0))
	fmt.Printf("\n the Results are: %v",results)
}


func dbCall(i int){
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from database is:",dbData[i])

	// using mutex
	// a.Lock()  // check to see if locked 
	// results = append(results,dbData[i])
	// a.Unlock()  // check to release lock
	
	// using rw mutex
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string){
	a.Lock()
	results = append(results,result)
	a.Unlock()
}

func log(){
	a.RLock()
	fmt.Printf("\n the current results are: %v",results)
	a.RUnlock()
}

// A sync.RWMutex allows multiple goroutines to read a resource simultaneously but ensures that only one can write at a time.

// Key Methods:
// RLock(): Acquires a read lock. Multiple goroutines can hold read locks simultaneously.
// RUnlock(): Releases a read lock.
// Lock(): Acquires a write lock. No other goroutine can read or write while a write lock is held.
// Unlock(): Releases the write lock.


