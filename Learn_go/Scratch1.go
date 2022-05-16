package main

import "fmt"

var (
	age int = 19
	name string ="Luffy"
	destined string = "King of the pirates"
)
func main(){
	fmt.Println("orewa ", name,age, destined,"orewanaru");

	for i:= 2; i<= 10; i++{
		if (i%2==0){
			fmt.Println("even digit ",i)
		}

	}
	cel:= 100
	fah:= (cel * 9/5) + 32
	fmt.Println("the value of fah:", fah)


	
}
