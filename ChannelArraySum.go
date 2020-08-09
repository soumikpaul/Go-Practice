package main

import ("fmt"
  "time"
)

func sum(arr []int, c chan int)  {
  sum := 0
	for _, v := range arr {
		sum += v
	}
  // time.Sleep(time.Second)
	c <- sum
}


func main()  {
    arr := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
    c := make(chan int)
    start := time.Now()
    go sum(arr[:len(arr)/2],c)
    go sum(arr[len(arr)/2:],c)

    // go sum(arr,c)


    fmt.Println(<-c + <-c)
    // fmt.Println(<-c)

    end:= time.Now()
    fmt.Println(end,start)
    fmt.Println(end.Sub(start))

}
