package main

import (
  "fmt"
  "reflect"

)

func main(){
  languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust", // Must include the trailing comma
	}

  // -- snip -- //
        allLangs := languages[:]                      // copy of the array
        fmt.Println(reflect.TypeOf(allLangs).Kind())   // slice
  
  /* Create a slice containing web frameworks */
        frameworks := []string{
            "React", "Vue", "Angular", "Svelte",
            "Laravel", "Django", "Flask", "Fiber",
        }

  jsFramework := frameworks[0:4:4]
  frameworks = append(frameworks,"Meteor")

  fmt.Println("all frameworks : %v \n" , frameworks )
  fmt.Println("javascript fm : %v \n", jsFramework)


}
