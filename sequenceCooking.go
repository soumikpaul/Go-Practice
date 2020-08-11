package main

import ("fmt"
  "math/rand"
    "time")


type CookInfo struct {
  foodCooked  string
  waitForPartner chan bool
}
func cookFood(name string)<- chan CookInfo  {
  cookChannel := make(chan CookInfo)
	wait := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			cookChannel<- CookInfo{fmt.Sprintf("%s %s", name,"Done") , wait}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			<-wait
		}
	}()

	return cookChannel
}
func fanIn(chan1, chan2 <-chan CookInfo) <-chan CookInfo  {

  myChannel := make(chan CookInfo)

  go func ()  {
    for {
      myChannel <- <-chan1
    }
  }()

  go func ()  {
    for {
      myChannel <- <-chan2
    }
  }()

  return myChannel
}
func main()  {
  finalCh := fanIn(cookFood("Soumik"),cookFood("Sanu"))

  for i:=0; i<3; i++ {
    food1 := <- finalCh
    fmt.Println(food1.foodCooked)

    food2 := <- finalCh
    fmt.Println(food2.foodCooked)

    food1.waitForPartner <- true
    food2.waitForPartner <- true


    fmt.Printf("Done with round %d\n", i+1)

  }
}
