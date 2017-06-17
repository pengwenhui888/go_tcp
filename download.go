package main
import (
    "fmt"
    "time"
)
func main() {
    const (
        GOROUTINE_COUNT = 2
        TASK_COUNT      = 10
    )
    chReq := make(chan string, GOROUTINE_COUNT)
    chRes := make(chan int, GOROUTINE_COUNT)
    for i := 0; i < GOROUTINE_COUNT; i++ {
        go func() {
            for {
                url := <-chReq

                fmt.Println(url)
                time.Sleep(time.Second * 2)
                chRes <- 0
            }
        }()
    }
    urls := make([]string, TASK_COUNT)
    for i := 0; i < TASK_COUNT; i++ {
        urls[i] = fmt.Sprintf("http://www.%d.com", i)
    }
    go func() {
        // got urls
        for i := 0; i < TASK_COUNT; i++ {
            chReq <- urls[i]
        }
    }()
    for i := 0; i < TASK_COUNT; i++ {
       <-chRes
        // check error
    }
}
