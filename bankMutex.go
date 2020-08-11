package main

import ("fmt"
    "sync")

func deposit(balance *int, val int, wg *sync.WaitGroup, mutex *sync.Mutex)  {
  mutex.Lock()
  *balance+=val
  mutex.Unlock()
  wg.Done()
}

func withdraw(balance *int, val int, wg *sync.WaitGroup, mutex *sync.Mutex)  {
  mutex.Lock()
  *balance-=val
  mutex.Unlock()
  wg.Done()
}


func main()  {

  var wg sync.WaitGroup
  var mutex sync.Mutex

  balance:= 100

  wg.Add(2)
  go deposit(&balance, 50, &wg, &mutex)
  go withdraw(&balance, 30, &wg, &mutex)
  // withdraw(&balance, 30, &wg, &mutex)

  wg.Wait()
  fmt.Println(balance)


}
