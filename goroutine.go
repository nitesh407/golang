package main
import (
	"fmt"
	"time"
)

func Display(str string){
	for i:=1; i<6; i++{
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

func main(){

	//goroutine call
	go Display("This is the Goroutine call")

	//Normal call
	Display("This is the Normal call")
}