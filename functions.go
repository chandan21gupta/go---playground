package main

import "fmt"

func add(x int, y int) int {
  return x + y
}


func two_items(x string, y int) (string, int) {
  return x, y
}

func main() {
  fmt.Println(two_items("42",13))
}
