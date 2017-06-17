package main

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "log"
  "os"
  "time"
)

func main (){
  var param []interface{}
//  datetime := time.Now()
  datetime := time.Now().Format("2006-01-02 15:04:05")

  numbers := [3]string{"pony","bbbs",datetime}
  fmt.Println(datetime)
  for i,_ := range numbers {
    param = append(param,numbers[i])
  }

  hosts := "root:12345678@tcp(127.0.0.1:3306)/passport?charset=utf8"
  db, _ := sql.Open("mysql", hosts)
  stmt, err := db.Prepare("INSERT user(username,password,addtime) values(?,?,?)")
  checkErr(err)
  res, err := stmt.Exec(param...)
  checkErr(err)
  ids, err := res.LastInsertId()
  checkErr(err)
  fmt.Println(ids)
  fmt.Println(param)

  sqls := "select id,username from user"
  data := readRows(db,sqls)
  var id int
  var username string
  for data.Next() {
    err := data.Scan(&id, &username)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(id, username)
  }
}

func readRows(db *sql.DB,sqlstring string) *sql.Rows {
  rows, err := db.Query(sqlstring)
  if err != nil {
    fmt.Println("fetech data failed:", err.Error())
  }
  //defer rows.Close()
  return rows
}

func LogErr(v ...interface{}) {

	logfile := os.Stdout
	log.Println(v...)
	logger := log.New(logfile,"\r\n",log.Llongfile|log.Ldate|log.Ltime);
	logger.SetPrefix("[Error]")
	logger.Println(v...)
	defer logfile.Close();
}

func checkErr(err error) {
	if err != nil {
		LogErr(os.Stderr, "Fatal error: %s", err.Error())
	}
}
