package main
import ("fmt"
  "time"
)

func main()  {
  myMsg := make(chan string )

  go func ()  {
    myMsg <- "hello"
    myMsg <- "world"
  }()

  resMsg := <- myMsg
  resMsg2 := <- myMsg
  fmt.Println(resMsg)
  fmt.Println(resMsg2)

  anotherChannel := make(chan string,2)

  anotherChannel <- "soumik"
  anotherChannel <- "zeba"

  fmt.Println(<-anotherChannel)
  fmt.Println(<-anotherChannel)

// This is channel synchronization
// used for making one routine wait for another
  done := make(chan bool, 1)
  go worker(done)

  <-done

}
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    time.Sleep(time.Second)

    fmt.Println("done")
    done <- true
  }
