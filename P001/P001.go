package main

import "fmt"

func main(){

  sum := 0

  for i := 3; i < 1000; i++ {
    if i%5 == 0 || i%3 == 0 {
      sum = sum + i
    }
  }

  fmt.Printf("%d\n", sum)
  

}