package main

import ( "fmt"
        "time")



func main() {
  channel1 := make(chan string)
  channel2 := make(chan string)

  go func()  {
    for i:= 0;i<5;i++{
      channel1 <- "Soumik"
      time.Sleep(time.Millisecond * 100)

    }
  } ()

  go func ()  {
    for i:= 0; i<5; i++{
      channel2 <- "Sanu"
      time.Sleep(time.Millisecond * 100)

    }
  } ()

  for i:= 0; i<5; i++ {
      select {
      case message1 := <-channel1:
        fmt.Println(message1)

      case message2 := <-channel2:
        fmt.Println(message2)

      default :
        fmt.Println("zeba")
      }


    }

}
