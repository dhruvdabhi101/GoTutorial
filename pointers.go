package main

import "fmt"


func main(){
  var address *int
  num := 42
  address = &num
  value := *address

  fmt.Println(address)
  fmt.Println(value)

}
