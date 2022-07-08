package main

import "fmt"

type stack struct{
  index int
  data [10]int

}

func (s *stack) push(k int){
  s.data[s.index] = k
  s.index++
}
func (s *stack) pop()int {
  num := s.data[s.index]
  s.index--
  return num

}
func main(){
  s := new(stack) // new keyword for creating struct objects
  s.push(12)
  s.push(21)
  fmt.Println(*s)


}
