package main

// funtcion closures

import "fmt"



type person struct {
  name string
  age int
}

func main() {
    fmt.Println("hello world")
        // fmt.Println(d)
    for i:=0;i<10;i++{
      fmt.Println(i)
    }
    left:=5
    fmt.Println("Res",left>>1)

    var arr[5] int
    arr[0]=9
    arr[1]=1

    for i:=0;i<len(arr);i++{
      fmt.Println("val ",arr[i])
      }

    b:= [5]int {1,2,3,4,5}
    fmt.Println("b: ",b)
    sum:=0
    for iter,number:= range b{
      sum+=number
      fmt.Println("iter",iter)
    }
    fmt.Println(sum)

    x := "soumik"
    fmt.Println(x)
    // x[0]="z"
    fmt.Println(x)
    s:=make([]string, 3)

    s[0]="a"
    s[1]="soumik"
    s[2]="a"
    fmt.Println(s)
    s[0]="z"
    s = append(s,"zeb")
    fmt.Println(s)
    fmt.Println(add(2,3))
    add,sub,mul := multiVal(5,2)
    fmt.Println(mul)
    fmt.Println("sub",sub," add ",add)
    multiArg(1,2,3,4)

    intSeq := closeIntSeq()
    fmt.Println(intSeq())
    fmt.Println(intSeq())
    fmt.Println(intSeq(),intSeq())
    fmt.Println("fact = ",factorial(5))
    l := 2
    zeroPtr(&l)
    fmt.Println(" l = ",l)
    fmt.Println(person{"soumik",25})
    newP := person{name:"sanu",age:26}
    fmt.Println(newP)
    newP.age = 36
    fmt.Println(newP)
}
func add(a int,b int) int  {
  return a+b
}

func multiVal(a int,b int) (int, int, int)  {
  return a+b,a-b,a*b
}
func multiArg(arr ...int){
  sum:=0
  for _, num := range arr{
    sum+=num
  }
fmt.Println(sum)
}
func closeIntSeq() func () int  {
  i:=0
  return func () int {
    i++
    return i
  }
}
func factorial(a int) int {
  if a==0 {
    return 1
  }
  return a*factorial(a-1)
}
func zeroPtr(a *int){
  fmt.Println(a)
  *a = 10
}
