package main
import "fmt"

func main() {
  var input int
  fmt.Scanln(&input)
  for input < 100 {
    input+=input
    fmt.Println(input)
  }
}
