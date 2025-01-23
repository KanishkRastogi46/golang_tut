package main

import (
	"fmt"
	"sync"
)

func task (id int, wg *sync.WaitGroup) {
	// it will be called when the function ends and removes the task from the wait group
	defer wg.Done()
	fmt.Println("dont task", id)
}

func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		// it adds task to the wait group
		wg.Add(1)
		go task(i, &wg)
	}
	// it will wait until all the tasks are done
	wg.Wait()
}