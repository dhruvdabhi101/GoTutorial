package main

import "fmt"


func main(){
  release := map[string]int{
    "dhruv":2003, "nachiket":2002, "rushi":2003, "Python":1991,
  }
  for k , v := range release{
    fmt.Printf("name %s and year %d \n",k,v)
  }


}
