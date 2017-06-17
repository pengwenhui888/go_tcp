package main


import (
    "fmt"
    "reflect"
)


func main() {
   tonydon := &User{"pwh", 100, "0000123"}
   object := reflect.ValueOf(tonydon)
   fmt.Println(object)
   myref := object.Elem()
      fmt.Println(myref)
   typeOfType := myref.Type()
     fmt.Println(typeOfType)
      fmt.Println(myref.NumField())
   for i:=0; i<myref.NumField(); i++{
       field := myref.Field(i)
       fmt.Printf("%d. %s %s = %v \n", i, typeOfType.Field(i).Name, field.Type(), field.Interface())
   }
   tonydon.SayHello()
  // v := object.MethodByName("SayHello")
   //v.Call([]reflect.Value{})
}

type User struct {
    Name string
    Age  int
    Id   string
}


func (u *User) SayHello() {
    fmt.Println("I'm " + u.Name + ", Id is " + u.Id + ". Nice to meet you! ")
}
