package main

import "fmt"

func main(){
  languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust", // Must include the trailing comma
	}

  classic := languages[0:3] // alt -> [:3]
  modern := make([]string,4) // len(modern) = 4 
  modern = languages[3:7] // include 3 exclude 7 
  new := languages[7:9] // alt [7:]

  fmt.Println("Classic Languages : ", classic)
  fmt.Println("Modern Languages", modern)
  fmt.Println("New Languages", new)


}
