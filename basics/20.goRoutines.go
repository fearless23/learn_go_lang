package basics

import (
	"sync"
	"time"
)

// ####  BASIC  ####

func Goroutines1() {
	msg := "hello"
	go printMsg1(msg) // runs in different thread and takes 3 seconds
	msg = "world"
	// main will exist before printMsg1 finishes
	// So, not printMsg
}

// Function takes 3 seconds to complete
func printMsg1(msg string) {
	time.Sleep(3 * time.Second)
	Log("msg", msg)
}

// ####  WITH WaitGroup  ####
var wg = sync.WaitGroup{}

func Goroutines2() {
	msg := "hello"
	wg.Add(1)
	go printMsg2(msg) // runs in different thread
	msg = "world"

	Log("Waiting for main to finsh; msg", msg)
	wg.Wait() // wait for all goroutine to finish
}

// Function takes 3 seconds to complete
func printMsg2(msg string) {
	time.Sleep(3 * time.Second)
	Log("msg", msg)
	wg.Done() // This is necessary to tell waitGroup if a goroutine has finished
}

// ####  WITH WaitGroup and more  ####
var counter = 0

func Goroutines3() {
	msg := "hello"
	for i := 0; i < 3; i++ {
		wg.Add(2) // add 2 goroutines (mismatch in number of goroutine will throw error)
		go sayHello(msg, i)
		go increment()
	}

	msg = "world"
	Log("Waiting for main to finsh; msg", msg)
	wg.Wait()
}

func sayHello(msg string, i int) {
	Log("msg-"+intToStr(i), msg) // message is always hello
	// time.Sleep(1 * time.Second)
	Log("counter "+intToStr(i), counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}

// 6 goroutines are started one after another
// 1. i=0 sayHello   PRINTS - (msg=Hello,i=0,counter=globalValue)
// 2. i=0 increment  (counter= globalValue++)
// 3. i=0 sayHello   PRINTS - (msg=Hello,i=0,counter=globalValue)
// 4. i=0 increment  (counter= globalValue++)
// 5. i=0 sayHello   PRINTS - (msg=Hello,i=0,counter=globalValue)
// 6. i=0 increment  (counter= globalValue++)

// by the time we reach the line: `Log("counter", counter)`, counter value can be anything from 0 to 3
// and any of 6 goroutines can finish before each other.
