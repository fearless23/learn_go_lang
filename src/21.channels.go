package src

import (
	"fmt"
	"sync"
	"time"
)

var wag = sync.WaitGroup{}

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

var logChannel = make(chan logEntry, 50)

func MyChannels() {
	// ch := make(chan int)
	// gologl(ch)
	// runTimes(ch, 2)

	defer closeLogger()
	go myLogger()
	logChannel <- logEntry{time.Now(), logInfo, "App is Starting"}
	logChannel <- logEntry{time.Now(), logError, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	// wag.Wait()
}

/*
Channel is a way to share data between go
routine.
*/

// Receive Data to Channel
func receive(ch <-chan int) {
	i := <-ch
	print(i)
	wag.Done()
}

// Send some data to Channel
func send(ch chan<- int, i int) {
	ch <- i
	wag.Done()
}

func runTimes(ch chan int, n int) {
	for i := 0; i < n; i++ {
		wag.Add(2)
		go receive(ch)
		go send(ch, i)
	}
}

func myLogger() {
	for entry := range logChannel {
		fmt.Printf("%v - [%v] %v\n", entry.time.Format("Mon 02 January 2006 03:04:05 PM"), entry.severity, entry.msg)
	}
}

func closeLogger() {
	close(logChannel)
}
