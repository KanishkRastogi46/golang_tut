package main

import (
	"fmt"
)

// func changeNum (num chan int) {
// 	for n := range num {
// 		fmt.Println("processing number:", n)
// 		time.Sleep(time.Second)
// 	}
// }

// func sum (res chan int, a int, b int) {
// 	res <- a + b
// }

// func decision (accept chan bool) {
// 	defer func () { accept <- true } ()
// 	fmt.Println("Processing...")
// }

func createEmail (emailChan chan string) {
	for email := range emailChan {
		fmt.Println("Sending email to:", email)
	}
}

func main() {
	// messageChannel := make(chan string)
	// messageChannel <- "Hello, World!"
	// message := <-messageChannel
	// fmt.Println(message)

	// num := make(chan int)
	// go changeNum(num)
	// for {
	// 	num <- rand.Intn(100)
	// }
	// time.Sleep(time.Second*2)

	// result := make(chan int)
	// go sum(result, 1, 2)
	// res := <-result		// blocking
	// fmt.Println("Result:", res)

	// accept := make(chan bool)
	// go decision(accept)
	// <- accept			// blocking

	emailChan := make(chan string, 20)
	go createEmail(emailChan) 
	emailChan <- "email1"
	emailChan <- "email2"
}