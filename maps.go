package main
import "fmt"

func main(){

	// we also gonna see the difference bw the make and the new keyword
	/*
	var score map[string]int
	score["Nitesh"] = 86
	fmt.Println(score)           //panic: assignment to entry in nil map
	*/

	score := make(map[string]int)
	score["Nitesh"] = 86
	score["Bharat"] = 89
	score["Pawan"] = 23
	score["Madhav"] = 24
	score["mohit"] = 46
	fmt.Println(score)

	//we can also loop through map, it wont print in order
	for k, v := range score{
		fmt.Printf("The key is %v and the value is %v\n", k, v)
	}
	//we can delete as well
	delete(score, "Bharat")
	fmt.Println(score)

	age := map[string]int{ "Nitesh":21, "Baman": 23}
	fmt.Println(age)

	val, _ := score["Nitesh"]   //this returns the value and the boolean value
	fmt.Println(val)
}