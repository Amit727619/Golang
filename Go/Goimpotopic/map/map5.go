package main
import ("fmt")

func main() {
  var a = map[string]string{"name": "durvesh", "addr": "pune", "year": "2000"}
  b := a

  fmt.Println(a)
  fmt.Println(b)

  b["year"] = "2020"
  fmt.Println("After change to b:")

  fmt.Println(a)
  fmt.Println(b)
}