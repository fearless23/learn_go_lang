package main

import (
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func goroutines() {
	msg := "hello"

	// wg.Add(1)
	// go sayHello(msg)
	// msg = "no,hello"
	// time.Sleep(100 * time.Millisecond)

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello(msg)
		m.Lock()
		go increment()
	}

	wg.Wait()

	// runtime.GOMAXPROCS(100)
	// x := runtime.GOMAXPROCS(-1)

}

func sayHello(msg string) {
	gologv("Counter", counter)
	gologv("msg", msg)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
