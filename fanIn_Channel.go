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

//For fanIn process we need more than one producer to send data to channel and one consumer to read data from data from channel
func Producer(ch chan <- int, name string){
	for{
		//sleep for some time
		sleep()
		//generate a random number
		n := rand.Intn(100)
		//send the message to channel
		fmt.Printf("Channel %s -> %d\n", name, n)
		ch <- n
	}
}

func Consumer(ch <- chan int){
	for n:= range ch{
		fmt.Printf("<- %d\n", n)
	}
}

func FanIn(chA, chB <- chan int, chC chan <- int){
	var n int
	for{
		select{
		case n = <- chA:
			chC <- n
		case n = <- chB:
			chC <- n
		}
	}
}

func  main(){
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go Producer(chA, "chA")
	go Producer(chB, "chB")
	go Consumer(chC)

	// now we create the fanin function to actually read data from Producers and write it to the consumer
	FanIn(chA, chB, chC)
}