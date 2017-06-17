package db

import(
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "encoding/json"
  "io/ioutil"
  "strings"
  "log"
  "os"
)

var db *sql.DB

type Config struct {
  Host string `json:"host"`
  Password string  `json:"password"`
  Username string `json:"username"`
  Databasename string `json:"databasename"`
}


func (c *Config)Run(){
  host := c.Host
  password := c.Password
  username := c.Username
  databasename := c.Databasename
  hosts := username+":"+password+"@tcp("+host+")/"+databasename+"?charset=utf8"
  db, _ = sql.Open("mysql", hosts)
  db.SetMaxOpenConns(3)
  db.SetMaxIdleConns(1)
}

func init(){
  var conf Config
  bytes, err := ioutil.ReadFile("/home/go/config/database.json")
  if err != nil {
      fmt.Println("ReadFile: ", err.Error())
  }

  if err := json.Unmarshal(bytes, &conf); err != nil {
      fmt.Println("Unmarshal: ", err.Error())
  }
  conf.Run()
}

func ReadConfig() (map[string]string, error) {
   var config = map[string]string{}
    bytes, err := ioutil.ReadFile("/home/go/config/database.json")
    if err != nil {
        fmt.Println("ReadFile: ", err.Error())
        return nil, err
    }
    if err := json.Unmarshal(bytes, &config); err != nil {
        fmt.Println("Unmarshal: ", err.Error())
        return nil, err
    }

    return config, nil
  }

func Query(sqlstring string) string{
    rows, err := db.Query(sqlstring)
    if err != nil {
      fmt.Println("fetech data failed:", err.Error())
      return ""
    }
    defer rows.Close()
    json := getJson(rows)
    return json
    //fmt.Println(string(jsonData))
}

func getJson(rows *sql.Rows) string{
  columns, _ := rows.Columns()
  count := len(columns)
  tableData := make([]map[string]interface{}, 0)
  values := make([]interface{}, count)
  valuePtrs := make([]interface{}, count)
  for rows.Next() {
    for i := 0; i < count; i++ {
      valuePtrs[i] = &values[i]
    }
    rows.Scan(valuePtrs...)
    entry := make(map[string]interface{})
    for i, col := range columns {
      var v interface{}
      val := values[i]
      b, ok := val.([]byte)
      if ok {
        v = string(b)
      } else {
        v = val
      }
      entry[col] = v
    }
    tableData = append(tableData, entry)
  }
  jsonData, err := json.Marshal(tableData)
  if err != nil {
    fmt.Println(err)
  }
  return string(jsonData)
}


func readRows(sqlstring string) *sql.Rows{
  rows, err := db.Query(sqlstring)
  if err != nil {
    fmt.Println("fetech data failed:", err.Error())
  }
  defer rows.Close()
  return rows
}

func Scheduler(sql string) string{
    if strings.Contains(sql,"select") {

    }else {

    }
    return ""
}

func insert(sql string,param []interface{}) int64 {
	stmt, err := db.Prepare(sql)
	checkErr(err)
	res, err := stmt.Exec(param...)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
  return id
}

func update(sql string,param []interface{}) int64{
	stmt, err := db.Prepare(sql)
	checkErr(err)
	res, err := stmt.Exec(param...)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	return num
}

func remove(sql string,param []interface{}) int64{
	stmt, err := db.Prepare(sql)
	checkErr(err)
	res, err := stmt.Exec(param...)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
  return num
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
