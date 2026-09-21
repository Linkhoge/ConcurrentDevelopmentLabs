package main

import (
	"fmt"
	"sync"
	"time"
)

// --------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Ariel Fajimiyo C00300811
// c00300811@setu.ie
// 21/09/2026
// -------------------------------------------

// make struct containing channel
// add init, acquire and release
type semaphore struct {
	theCounter chan struct{}
}

// funcAcquite(sem *Semaphore)
func main() {
	maxGoroutines := 5
	semaphore := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// Simulate a task
			fmt.Printf("Running task %d\n", i)
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}
