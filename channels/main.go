package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		defer close(ch)
		ch <- "Hola channel"
	}()
	fmt.Println(<-ch)

	// ch1 := make(chan int)
	// go func() {
	// 	defer close(ch1)
	// 	ch1 <- 1
	// 	ch1 <- 2
	// 	ch1 <- 3
	// 	ch1 <- 4
	// }()

	// for n := range ch1 {
	// 	fmt.Printf("%d\n", n)
	// }

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	// ch2 <- 3 ESTO LEVANTARÍA UN ERROR

	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	// Cada que yo envío información, libero un espacio
	ch2 <- 3
	ch2 <- 4
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

}
