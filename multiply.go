package main

import ("fmt"
    "sync"
    "time")

func printTable(n int, wg *sync.WaitGroup)  {
  for i := 1; i <= 12; i++ {
    fmt.Printf("%d x %d = %d\n", i, n, n*i)
    // time.Sleep(time.Millisecond * 50)

  }
  wg.Done()
}

func main() {
  var wg sync.WaitGroup
  // wg.Add(12)
  for i:=0; i<12;i++ {
    wg.Add(1)
    go printTable(i,&wg)
    time.Sleep(time.Millisecond * 50)

  }
  wg.Wait()
}
