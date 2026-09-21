package main

import (
	"fmt"
	"sync"
)

// --------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Ariel Fajimiyo C00300811
// c00300811@setu.ie
// 21/09/2026
// -------------------------------------------

// Global variables shared between functions --A BAD IDEA
var wg sync.WaitGroup
var total int64

func adds(n int, theLock *sync.Mutex) bool {
	for i := 0; i < n; i++ {
		// lock the mutex
		theLock.Lock()
		total++
		// unlock the mutex
		theLock.Unlock()
	}
	wg.Done() //let waitgroup know we have finished
	return true
}

func main() {

	//theLock will be passed by reference between go routines
	//better than using a global variable
	var theLock sync.Mutex

	total = 0
	//the waitgroup is used as a barrier
	// init it to number of go routines
	wg.Add(10)

	//for loop using range option
	for i := range 10 {
		//starting
		fmt.Println(i)
		go adds(1000, &theLock)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println(total)
}
