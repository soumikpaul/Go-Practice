package main
import "fmt"

type rect struct {
  length int
  breadth int
}
func (r *rect) area() int  {
  return r.length*r.breadth
}


func main()  {
  r:=rect{breadth:5,length:10}
  fmt.Println(r.area())
}
