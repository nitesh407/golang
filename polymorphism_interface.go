package main

import (
	"fmt"
)

type emp interface{
	develop()int
	name()string
}

type team_1 struct{
	
	total_1 int
	name_1 string
}

func (t team_1)develop()int{

	return t.total_1
}

func (t team_1)name()string{

	return t.name_1
}

type team_2 struct{
	
	total_2 int
	name_2 string
}

func (t team_2)develop()int{

	return t.total_2
}

func (t team_2)name()string{

	return t.name_2
}

func final_develop(s []emp){

	total_project := 0
	for _, ele := range s{
		
		fmt.Println("Project enviroment: ", ele.name())
		fmt.Println("Total projects: ", ele.develop())
		total_project += ele.develop()
	}
	fmt.Println("Total projects developed by the company: ", total_project)
}

func main() {
	
	res1 := team_1{
			total_1 : 25,
			name_1 : "Golang",
	}
	
	res2 := team_2{
			total_2 : 25,
			name_2 : "C",
	}
	//fmt.Println(res1, res2)
	
	final := []emp{res1, res2}
	final_develop(final)
}	