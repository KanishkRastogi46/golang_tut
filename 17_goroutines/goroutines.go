package main

import (
	"fmt"
	"time"
)

func task (id int) {
	fmt.Println("dont task", id)
}

func main() {
	for i := range 10 {
		go task(i)
	}

	time.Sleep(time.Second * 2)
}