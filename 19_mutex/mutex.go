package main

import (
	"fmt"
	"sync"
)

type post struct {
	likes int
	mut sync.Mutex
}

func (p *post) increment(wg *sync.WaitGroup) {
	defer func () { 
		p.mut.Unlock()
		wg.Done()
	}()

	p.mut.Lock()
	p.likes = p.likes + 1
}

func main() {
	p := post{likes: 0}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go p.increment(&wg)
	}
	wg.Wait()
	fmt.Println(p.likes)
}