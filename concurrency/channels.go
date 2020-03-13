package main

import (
	"fmt"
)

func main() {
	// Unbuffered channel
	// ch := make(chan int)

	// go func() {
	// 	ch <- 10
	// 	fmt.Println("sent")
	// }()

	// fmt.Println(<-ch) // láy trị từ goroutines ra nếu ko thì sẽ bị block
	// fmt.Println("done")

	// buffered channel
	// ch := make(chan int, 2) // nếu là 0 chính là unbuffered channel, 2 là capacity
	// ch <- 5
	// ch <- 10

	// close(ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch) // nếu có close() mà dư capacity nó trả về là 0
	// close channel (case 2)
	queue := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			if i > 5 {
				close(queue)
				break
			} else {
				queue <- i
			}
		}
	}()

	for value := range queue {
		fmt.Println(value)
	}
	// select
	// queue := make(chan int)
	// done := make(chan bool)
	// fmt.Println(1)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		queue <- i
	// 	}
	// 	done <- true
	// }()

	// for {
	// 	select {
	// 	case v := <-queue:
	// 		fmt.Println(v)
	// 	case <-done:
	// 		fmt.Println("done")
	// 		return
	// 	}

	// }
}
