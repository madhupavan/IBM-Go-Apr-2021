package main

import (
	"fmt"
	"sync"
	"time"
)

func goRoutine1(wg *sync.WaitGroup) {
	//wg.Add(1)
	fmt.Println("starting goRoutine1")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("goRoutine1 completed")
	wg.Done()
}

func goRoutine2(wg *sync.WaitGroup) {
	//wg.Add(1)
	fmt.Println("starting goRoutine2")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("goRoutine2 completed")
	wg.Done()
}

func goRoutine3(wg *sync.WaitGroup) {
	//wg.Add(1)
	fmt.Println("starting goRoutine3")
	time.Sleep(750 * time.Millisecond)
	fmt.Println("goRoutine3 completed")
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go goRoutine1(wg)
	go goRoutine2(wg)
	go goRoutine3(wg)
	wg.Wait()
	fmt.Println("Done..!")
}
