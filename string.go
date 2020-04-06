package main
import (
	"fmt"
	"strings"
)

func main(){

	str := "Nitesh"

	var b = make([]byte, len(str))
	b = []byte(str)
	fmt.Println(b)

	s := "Kumar"
	fmt.Println(strings.Compare(s, str))
	//concatination of sltrings
	fmt.Println("Using + operator :", str+s+ "Sharma")

	//using fmt.Sprintf() method
	str2 := fmt.Sprintf("%s%s", str, s)
	fmt.Println(str2)

	//Joining a slice of string
	slice := []string{"Nitesh", "Kumar", "Sharma"}
	slice_string := strings.Join(slice, " ")
	fmt.Println(slice_string)

	fmt.Println("Trimmed string : ", strings.Trim(slice_string, "a"))

	strr1 := "Welcome, to the, online portal, of GeeksforGeeks"
    strr2 := "My dog name is Dollar"

    fmt.Println("\nResult 1: ", strings.Split(strr1, ",")) 
    fmt.Println("Result 2: ", strings.Split(strr2, ""))
	

	//we have converted the string to a slice of string    
    ss := (strings.Split(str, ""))
    //And then we onverted the slice of string to string again	
    fmt.Println(strings.Join(ss, ""))

    //contains()
    fmt.Println(strings.Contains(slice_string, "Kumar"))
    //Repeat()
    fmt.Println(strings.Repeat(str, 3))
    //Index()
    fmt.Println(strings.Index(str, "s"))
    //Count()
    fmt.Println(strings.Count(slice_string, "a"))
    
}