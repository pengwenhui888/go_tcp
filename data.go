package main

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
 "github.com/go-redis/redis"
 "strconv"
)


var client *redis.Client

var db *sql.DB

func main()  {
  //连接MYSQL
  hosts := "root:1234567@tcp(127.0.0.1:3306)/passport?charset=utf8"
  db, _ = sql.Open("mysql", hosts)

  //连接REIDS
  //client = redis.New()
  //_ = client.Connect("511ad9aaf37c4d7a.m.cnqda.kvstore.aliyuncs.com", 6379)
  //client.Auth("asdQWE123")
  client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123", // no password set
		DB:       0,  // use default DB
	})
  taskCount := 4
  ch := make(chan int, taskCount)
  for i:=0; i < taskCount; i++{
      go func(i int){
        task(i)
        ch <- i
      }(i)
  }
  for i:=0; i < taskCount; i++{
        fmt.Println(<-ch)
  }

}

func task(i int){
    offset := i*10000
    s := "SELECT nickname from `user` where nickname <> '' order by id asc limit "+ strconv.Itoa(offset) + ",10000"
    rows, eq := db.Query(s)
    if eq != nil {
      fmt.Println(eq)
      return
    }
    defer rows.Close();
    k := 0
    for rows.Next(){
      var nickname string
      err  := rows.Scan(&nickname)
      if err != nil {
        fmt.Println(err)
        return
      }
      if (nickname != "") {
        client.SAdd("nickname",nickname)
        k++
      }
    }
    fmt.Println(k)
    fmt.Println(s)
}
