package main

import (

  //"database/sql"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "runtime"
  "math/rand"
  "strconv"
  "reflect"
)


type Config struct {
  Host string `json:"host"`
  Password string  `json:"password"`
  Username string `json:"username"`
  Databasename string `json:"databasename"`
}

type Tc struct {
  Td string
  Op string
  Port int
}

func (c * Config)Run(){
  host := c.Host
  password := c.Password
  username := c.Username
  databasename := c.Databasename
  hosts := username+":"+password+"@tcp("+host+")/"+databasename+"?charset=utf8"
  fmt.Println(hosts)
  //db, _ = sql.Open("mysql", hosts)
  //db.SetMaxOpenConns(3)
  //db.SetMaxIdleConns(1)
}

var complete chan int = make(chan int)

func loop() {
    for i := 0; i < 10; i++ {
        runtime.Gosched()
        fmt.Printf("%d ", i)
    }

    complete <- 0 // 执行完毕了，发个消息
}


func main() {
  var c Tc

  opName := "create"
  c.Op = opName
  c.Td = "666"
  c.Port = 3308

  fmt.Printf("conf.Port=%d\n\n", c.Port)

  // 结构信息
  t := reflect.TypeOf(c)
  // 值信息
  vs := reflect.ValueOf(c)
  var a []interface{}
  for s := range vs{
    a = append(a,s)
  }
fmt.Println(vs)
fmt.Println(a)

    runtime.GOMAXPROCS(2)
    go loop()
    go loop()
    for i := 0; i < 2; i++ {
      <- complete
    }

    var conf Config

    bytes, err := ioutil.ReadFile("/home/go/config/database.json")
    if err != nil {
        fmt.Println("ReadFile: ", err.Error())

    }

    if err := json.Unmarshal(bytes, &conf); err != nil {
        fmt.Println("Unmarshal: ", err.Error())

    }
    conf.Run()
    fmt.Printf("%v", conf)

    slice := make([]interface{}, 10)

    map1 := make(map[string]string)

    map2 := make(map[string]int)

    map2["TaskID"] = 1

    map1["Command"] = "ping"

    map3 := make(map[string]map[string]string)

    map3["mapvalue"] = map1

    slice[0] = map2

    slice[1] = map1

    slice[3] = map3

    fmt.Println(slice[0])

    fmt.Println(slice[1])

    fmt.Println(slice[3])
    limit :=  rand.Intn(5)
    fmt.Println(strconv.Itoa(limit))

}
