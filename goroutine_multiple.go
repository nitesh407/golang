package main

import (
	"fmt"
	"time"
)
func name(){

	arr := [4]string{"Nitesh", "Kumar", "Sharma", "Geek"}
	for i:=0; i<4; i++{
		time.Sleep(150 * time.Millisecond)
		fmt.Println("Name : ", arr[i])
	} 
}

func id(){
	
	arr:= [4]int{1, 2, 3, 4}
	for i:=0; i<4;i++{
		time.Sleep(500 * time.Millisecond)
		fmt.Println("id : ", arr[i])
	}
}

func main() {
	fmt.Println("Here is the main goroutine")
	
	go name()
	go id()
	time.Sleep(3500 * time.Millisecond) 
    	fmt.Println("\n!...Main Go-routine End...!") 
}