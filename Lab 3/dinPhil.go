// Dining Philosophers - channel solution using a footman
// Copyright (C) 2026 Ariel Fajimiyo
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2026
// Modified by: Ariel Fajimiyo
// Issues: Nan
// Students Helped:
//--------------------------------------------

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const philCount = 5
const mealCount = 3

// The left fork has the same number as the philosopher.
func left(index int) int {
	return index
}

// The right fork is the next fork around the table.
func right(index int) int {
	return (index + 1) % philCount
}

// Simulate a philosopher thinking for a random amount of time.
func think(index int) {
	x := time.Duration(rand.IntN(5))
	time.Sleep(x * time.Second)
	fmt.Printf("Phil %d was thinking\n", index)
}

// Simulate a philosopher eating for a random amount of time.
func eat(index int) {
	x := time.Duration(rand.IntN(5))
	time.Sleep(x * time.Second)
	fmt.Printf("Phil %d was eating\n", index)
}

// The footman limits entry to four philosophers before they take their forks.
func getForks(index int, footman chan bool, forks map[int]chan bool) {
	footman <- true
	forks[right(index)] <- true
	forks[left(index)] <- true
}

// Put down both forks, then leave the dining room by releasing the footman.
func putForks(index int, footman chan bool, forks map[int]chan bool) {
	<-forks[right(index)]
	<-forks[left(index)]
	<-footman
}

// Each philosopher thinks, eats, and releases the forks for each meal.
func doPhilStuff(index int, wg *sync.WaitGroup, footman chan bool, forks map[int]chan bool) {
	defer wg.Done()

	for meal := 0; meal < mealCount; meal++ {
		think(index)
		getForks(index, footman, forks)
		eat(index)
		putForks(index, footman, forks)
		fmt.Printf("Phil %d finished meal %d\n", index, meal+1)
	}
}

func main() {
	var wg sync.WaitGroup

	// Only four philosophers can try to pick up forks at the same time.
	footman := make(chan bool, philCount-1)

	// Each buffered channel represents one fork and allows only one holder.
	forks := make(map[int]chan bool)
	for i := 0; i < philCount; i++ {
		forks[i] = make(chan bool, 1)
	}

	wg.Add(philCount)
	for i := 0; i < philCount; i++ {
		go doPhilStuff(i, &wg, footman, forks)
	}

	wg.Wait()
	fmt.Println("All philosophers have finished.")
}
