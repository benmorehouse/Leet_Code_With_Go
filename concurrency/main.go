package main

import(
	"fmt"
	"time"
)

func timeToSleep(seconds int, x chan int){
	time.Sleep(time.Second * time.Duration(int64(seconds)))
	x <- seconds
}

func main(){
	var x chan int
	x = make(chan int)
	go timeToSleep(3, x)
	fmt.Println(<-x)
	close(x)
	c := make(chan bool, 4) // this is buffered which just means that this channel will have at most four values pushed throught it
	close(c)
}

