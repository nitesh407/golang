package main
import (
	"fmt"
	"math/rand"
	"time"
)

func echoworker(in <- chan int, out chan <- int){
	for{
		n := <-in
		time.Sleep(
			time.Duration(rand.Intn(3000)) * time.Millisecond,
		)

		out  <- n
	}
}

func Produce(ch chan <- int){
	i := 0
	for{
		fmt.Printf("-> Send Job : %d\n", i)
		ch <- i
		i++
	}
}
/*
func Consume(out <- chan int){

	for n:= range out{
		fmt.Printf("<- recv Job: %d\n", n)
	}
}
*/
func main(){
	in  := make(chan int)
	out := make(chan int)

	for i:=0; i<100; i++{
		go echoworker(in, out)
	}
	go Produce(in)
	//go Consume(out)

	for n:= range out{
		fmt.Printf("<- recv Job: %d\n", n)
	}

}