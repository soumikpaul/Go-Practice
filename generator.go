package main

import ("fmt")

func channelFromFunc() <-chan string {
  myChannel := make(chan string)
  go func ()  {
    for i:=0;;i++ {
      myChannel <- fmt.Sprintf("%s %d", "Counter at : ", i)
    }
  } ()
  return myChannel
}

func main()  {
  channel := channelFromFunc()

  for i:=0;i<5;i++ {
    fmt.Println(<-channel)
  }

  fmt.Println("Done with Counter")

}
