package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var dbData = []string{
	"id1",
	"id2",
	"id3",
	"id4",
	"id5",
}

// Wait groups let program wait for all goroutines to finish
var waitGroup = sync.WaitGroup{}
// Mutex can lock and unlock memory to prevent
var mutex = sync.Mutex{}

// Go provides a read-write mutex as well
var rwMutex = sync.RWMutex{}

var results = []string{}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is ", dbData[i])
	save(dbData[i])
	log()
	// This will signal that the goroutine has finished
	waitGroup.Done()
}

func main() {
	t0 := time.Now()
	for i := range len(dbData) {
		// This will signal that the goroutine has started
		waitGroup.Add(1)
		// dbCall(i)
		// To run concurrently use go keyword
		// This will not have program wait for dbCall to finish
		go dbCall(i)
	}
	// This will make the program wait for all dbCall to finish
	waitGroup.Wait()
	fmt.Printf("The program took %v to run.\n", time.Since(t0))
}

func save(result string) {
	// When things run concurrently, the results can be unexpected due to different processes writing to the memory
	// Mutex will help with this (Mutual Exclusion)
	mutex.Lock()
	results = append(results, result)
	mutex.Unlock()
}

func log() {
	// Can have many read locks but only one write lock
	// mutex locks will need to wait for read locks to finish
	rwMutex.RLock()
	fmt.Println("The results are ", results)
	rwMutex.RUnlock()
}