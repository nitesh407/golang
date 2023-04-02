//we cant sort a slice with the help of sort package directly, we need to convert it into a slice first	

package main
import "fmt"

func main(){
	var arr[3]int
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	fmt.Println(arr)

	a := [2]int{0}
	fmt.Println(a)

	//2D array

	arr1 := [3][3]int{
		{1, 2, 3}, 
		{4, 5, 6}, 
		{7, 8, 9},      // If we remove this line and try to acces these elements then it wont show the null pointer exception, 
						//because they are initialised by 0 implicitly
	}
	fmt.Println(arr1)
	fmt.Println(arr1[0])
	arr1[0][0] = 3

	//we we dont wanna mention the size of the array and want it to get the size as per the  number of elements given
	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(arr2)

	var a1 = make([]int, 5, 5)

	fmt.Println(len(a1))
	fmt.Println(cap(a))
	fmt.Println(len(a))


	array := [3]int{0, 1, 2}
	array_1 := array[:]        //we have created a slice here, so the changes will be done in the actual array

	array_1[0] = 10
	fmt.Println(array)

}