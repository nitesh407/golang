package main
import (
	"fmt"
	"bytes"
	"sort"
)

func main(){


	var s = make([]int, 3)
	s[0] = 0
	s[1] = 1
	s[2] = 2
	fmt.Println(s)

	//since we know that slices are of reference type. So, if we make
	//any changes to the slice, those changes will be showcased to the array
	a := [3]int{0, 1, 2}
	s1 := a[:]

	s1[0] = 5
	fmt.Println(a)
	if s1 == nil{
		fmt.Println("Slice is equal to nil")
	}else{
		fmt.Println("Not  equal to nil")
	}
	/*
	x := s==s1
	fmt.Println(x)
	*/// This will show error
	
	//Now cpoying a slice to another so that if we make any cnhanges to the newly created then those changes wont be reflected into previous array
	slice := make([]int, len(s))
	copy(slice, s)
	slice[0] = 100
	fmt.Println(s)	  //see the changes arent reflected in the base array

	//comparing slices using Compare()

	b := []byte{'N', 'i' ,'t', 'e', 's', 'h'}
	B := []byte{'n', 'i' ,'t', 'e', 's', 'h'}

	str  := string(b)
	fmt.Println("Str is ", str)
	fmt.Println(bytes.Compare(B, b))

	res := bytes.Equal(b, B)
	if res{
		fmt.Println("Equal")
	}else{
		fmt.Println("NE")
	}

	//sorting a slices
	sort.Ints(s1)
	fmt.Println(s1)
}