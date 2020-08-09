package main

import ("fmt"
"time"
)

func myFun(a string)  {
  for i:=0;i<5;i++ {
    fmt.Println(a," : ",i)
  }
}
func main()  {

  // myFun("normal")

  go myFun("routine")
  // go myFun("soumik")

  // fmt.Println("want to see printed in between or not")

  go func (a string)  {
    // fmt.Println("anonymous fn ",msg)
    for i:=0;i<5;i++ {
      fmt.Println(a," : ",i)
    }
  } ("passed")

  time.Sleep(time.Second)
  fmt.Println("done")




}
