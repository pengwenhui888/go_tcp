package main

import (
    "fmt"
    "net"
    "math/rand"
    "strconv"
    "time"
)

const (
    addr = "139.196.231.92:3333"
)

func main() {
    conn, err := net.Dial("tcp", addr)
    if err != nil {
        fmt.Println("server Connect error:", err.Error())
        return
    }
    fmt.Println("server Connect Success")
    defer conn.Close()
    Client(conn)
}

func Client(conn net.Conn) {
      //sms := make([]byte, 128)
    //for {
      //  fmt.Print("input send msg:")
      //  _, err := fmt.Scan(&sms)
      //  if err != nil {
      //      fmt.Println("data error:", err.Error())
      //  }
       rand.Seed(int64(time.Now().Nanosecond()))
        limit :=  strconv.Itoa(rand.Intn(50))
        sms := "select * from device_data limit " + limit
        conn.Write([]byte(sms))
        fmt.Println("sql send Success:",sms)
        buf := make([]byte, 10240)
        c, err := conn.Read(buf)
        if err != nil {
            fmt.Println("read server data error:", err.Error())
        }
        fmt.Println("sql return :",string(buf[0:c]))
    //}

}
