package main

import ("fmt"
   "math")

type geoInterface interface {
  area() float64
}
type circle struct {
  radius float64
}
// The method receiver appears in its own argument list between the func keyword and the method name.
 // Here is an example with a User struct containing two fields: FirstName and LastName of string type.


func (c circle) area() float64  {
  return math.Pi*c.radius*c.radius
}
func measure(g geoInterface)  {
  fmt.Println(g.area())
}

func main()  {
  c := circle{radius:5}
  measure(c)
}
