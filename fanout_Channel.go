package main
import (
	"fmt"
	"math/rand"
	"time"
)

func sleep(){
	time.Sleep(
		time.Duration(rand.Intn(3000)) * time.Millisecond,
	)
}

//For fanOut process we need one producer to send data to channel and more than one consumer to read data from data from channel
func Producer(ch chan <- int){
	for{
		//sleep for some time
		sleep()
		//generate a random number
		n := rand.Intn(100)
		//send the message to channel
		fmt.Printf("-> %d\n", n)
		ch <- n
	}
}

func Consumer(ch <- chan int, name string){
	for n:= range ch{
		fmt.Printf("Consumer %s <- %d\n",name,  n)
	}
}

func FanOut(chA <- chan int, chB, chC chan <- int){
	for n := range chA{
		if n <= 50{
			chB <- n
		}else{
			chC <- n
		}
	}

}

func  main(){
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go Producer(chA)
	go Consumer(chB, "chB")
	go Consumer(chC, "chC")

	// now we create the fanin function to actually read data from Producer and write it to the consumers
	FanOut(chA, chB, chC)
}