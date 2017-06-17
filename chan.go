package main

import(
  "fmt"
)

func main() {
    ch := make(chan int)
    for i:=0; i<5; i++ {
      go func(i int){
        //for{
            fmt.Println("Counting",i)
              ch <- i
        //}
      }(i)
    }

    for k:=0; k<5; k++ {
        fmt.Println(<-ch)
    }
}
