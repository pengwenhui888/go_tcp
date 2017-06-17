package main

import "fmt"

func main() {
	/*c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)*/


	ch1 := make(chan int, 1)

    ch2 := make(chan int, 1)

    ch2 <- 1

    select {

    case e1 := <-ch1:

        //如果ch1通道成功读取数据，则执行该case处理语句

        fmt.Printf("1th case is selected. e1=%v", e1)

    case e2 := <-ch2:

        //如果ch2通道成功读取数据，则执行该case处理语句

        fmt.Printf("2th case is selected. e2=%v", e2)

    default:

        //如果上面case都没有成功，则进入default处理流程

        fmt.Println("default!.")

    }
}

//select的轮询机制
func fibonacci(c chan int, quit chan int) {
	x, y := 2, 1
	for {
		select { // select轮询机制
		case c <- x:
			x = x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
