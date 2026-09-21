package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// --------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Ariel Fajimiyo C00300811
// c00300811@setu.ie
// 21/09/2026
// Students Helped: Adam Noonan
// -------------------------------------------

//Global variables shared between functions --A BAD IDEA

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, aArrived chan struct{}, bArrived chan struct{}) bool {

	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount

	fmt.Println("Part A", Num)

	//Rendezvous here
	aArrived <- struct{}{}

	// wait here until main says all threads have arrived
	<-bArrived

	fmt.Println("PartB", Num)

	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	threadCount := 3
	//barrier := make(chan bool)

	// create the channels for the rendezvous
	aArrived := make(chan struct{})
	bArrived := make(chan struct{})

	wg.Add(threadCount)

	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, aArrived, bArrived)
	}

	// wait here until all threads have arrived
	for range threadCount {
		<-aArrived
	}

	// tell all threads to continue
	for range threadCount {
		bArrived <- struct{}{}
	}

	wg.Wait() //wait here until everyone (10 go routines) is done

}
