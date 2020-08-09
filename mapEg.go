package main

import "fmt"

func main()  {
  myMap := map[string]int{"soumik":1,"zeba":2}
  fmt.Println(myMap)
  newMap := make(map[string]int)
  newMap["sanu"] = 5
  fmt.Println(newMap)
}
