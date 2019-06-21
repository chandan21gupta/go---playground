package main

import "fmt"

//var sum = 0

func main() {
  sum := 0
  for i:=0 ; i<=10; i++ {
    sum+=i
  }
  fmt.Println(sum)
}
