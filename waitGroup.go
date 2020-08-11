package main

import ("fmt"
"sync")

func welcomeMsg(wg *sync.WaitGroup){
  fmt.Println("Soumik from fn, 2nd one")
  wg.Done()
}

func main()  {

  var wg sync.WaitGroup

  wg.Add(2)

  go welcomeMsg(&wg)

  go func ()  {
    fmt.Println("Soumik first")
    wg.Done()
  }()

  wg.Wait()
}
