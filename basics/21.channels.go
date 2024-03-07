package basics

import (
	"fmt"
	"sync"
	"time"
)

// #### channel with  goroutines and waitGroup ####
var group1 = sync.WaitGroup{}

func Channels1() {
	ch := make(chan int) // create a channel with message type of int
	for i := 0; i < 2; i++ {
		group1.Add(2)
		go receive(ch, i) // start a goroutine and pass the channel
		go send(ch, i)    // start a goroutine and pass the channel
	}
	group1.Wait()
}

// Channels1 function will add 4 goroutines
// 1. i=0; receive  wait to read from channel
// 2. i=0; send     send the value 0 to channel
// 3. i=1; receive  wait to read from channel
// 4. i=1; send     send the value 1 to channel

// goroutine 1 & 3 are listeners and anyone can be first to listen
// goroutine 2 & 4 are senders and anyone can be first to send
// order of listening or sending depends upon thread spawing
// we can delay sending or receving using if block in next example

// Receive Data to Channel
func receive(ch <-chan int, i int) {
	val := <-ch // receives one data block from channel
	Log("reader:"+intToStr(i), val)
	group1.Done()
}

// Send some data to Channel
func send(ch chan<- int, i int) {
	Log("sender:"+intToStr(i), i)
	ch <- i // sends one data block to channel
	group1.Done()
}

// #### ----channel with  goroutines and waitGroup (example 2)---- ####
var group2 = sync.WaitGroup{}

func Channels2() {
	ch := make(chan int) // create a channel with message type of int
	for i := 0; i < 2; i++ {
		group2.Add(2)
		go receive2(ch, i) // start a goroutine and pass the channel
		go send2(ch, i)    // start a goroutine and pass the channel
	}
	group2.Wait()
}

// Channels2 function will add 4 goroutines
// 1. i=0; receive  sleep 1s; wait to read from channel
// 2. i=0; send     sleep 1s; send the value 0 to channel
// 3. i=1; receive  wait to read from channel
// 4. i=1; send     send the value 1 to channel

// goroutine 1 & 3 are listeners and since 1 sleep for 1s; receiver from goroutine3 is first listener
// goroutine 2 & 4 are senders and since goroutine2 sleep for 1s; sender from goroutine4 is first sender
// order of operations
// 1. [goroutine4] send value 1
// 2. [goroutine3] read value 1
// 3. [goroutine2] send value 0
// 4. [goroutine1] read value 0

// Receive Data to Channel
func receive2(ch <-chan int, i int) {
	if i == 0 {
		time.Sleep(1 * time.Second)
	}
	val := <-ch // receives one data block from channel
	Log("reader:"+intToStr(i), val)
	group2.Done()
}

// Send some data to Channel
func send2(ch chan<- int, i int) {
	if i == 0 {
		time.Sleep(1 * time.Second)
	}
	Log("sender:"+intToStr(i), i)
	ch <- i // sends one data block to channel
	group2.Done()
}

// #### ----(channel with defer and goroutine)---- ####

const (
	logInfo  = "INFO"
	logError = "ERROR"
	logWarn  = "WARN"
	logFatal = "FATAL"
)

type logEntry struct {
	time     time.Time
	severity string
	msg      string
}

var ch = make(chan logEntry)
var group3 = sync.WaitGroup{}

func Channels3() {
	// defer close(ch)
	group3.Add(1)
	go myLogger(ch)
	ch <- logEntry{time.Now(), logInfo, "App is Starting"}
	ch <- logEntry{time.Now(), logError, "App is shutting down"}
	group3.Wait()
}

func myLogger(ch <-chan logEntry) {
	for entry := range ch {
		fmt.Printf("%v - [%v] %v\n", entry.time.Format("Mon 02 January 2006 03:04:05 PM"), entry.severity, entry.msg)
	}
	group3.Done()
}

// Channels3
// in Channels2, send/receive functions read and write from pipe, and
// since 2 goroutines added to pipe and 2 goroutines took from pipe
// and once waitGroup for goroutines is done, channel is empty

// Whereas in Channels3 example; we have one goroutine which reads from channel; and
// main thread sends data to the channel

func closeLogger() {}
