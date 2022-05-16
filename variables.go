// Go lang is created by ROBERT GRIESEMAR, ROB PIKE AND KEN THOMPSON
//Go lang is Strong and Statically typed, static means variable should be declared in compile time

//Variables
// let us know how variable worls in go lang.

//Variable (var) variable_name variable_type
//Ex:  var i int

package main

import "fmt"

func main (){
	var i int
	i = 45
	fmt.Println("the values of i is :",i)

	ff()
}
var i,j int
func ff(){
	i = 15
	j = 12
	fmt.Println(i + j)

	gg()
}

var (

	first string 
	second float32
	age int
)

func gg() {

	first = "Monkey D Luffy"
	second = 1.5
	age = 17
	fmt.Println("orewa ",first," my bounty",second, " and my age is ",age)

	boo()

}

func boo() {
	//we can initialize and declare variable without metioning it's type, by using "colon"

	i:= 5
	fmt.Println("value is ",i)
}


