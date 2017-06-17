package main

import (
    "fmt"
    "net"
    "./db"
)

const (
    ip = "139.196.231.92"
    port = 3333
)



func main() {
    listen,err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), port, ""})
    if err != nil {
        fmt.Println("listen error:", err.Error())
        return
    }
    fmt.Println("waiting for listen ...")
    Server(listen)
}

func Server(listen *net.TCPListener) {
    for {
        conn, err := listen.AcceptTCP()
        if err != nil {
            fmt.Println("client error:", err.Error())
            continue
        }
        fmt.Println("client address:", conn.RemoteAddr().String())
        defer conn.Close()
        go func() {
            data := make([]byte, 128)
            //tmpBuffer := make([]byte, 0)
            //for {
            i, err := conn.Read(data)
            //tmpBuffer = utils.Depack(append(tmpBuffer, data[:i]...))
            list := List(string(data[:i]))
            fmt.Println("client data:", string(data[:i]))
            if err != nil {
                fmt.Println("read data error:", err.Error())
                //break
            }
            conn.Write([]byte(list))
            //}

        }()
    }
}


func List(sql string) string{
    data := db.Query(sql)
    return data
}
