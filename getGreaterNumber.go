package main

import(
	"fmt"
	"strconv"
)

func Printout(msg string){
	fmt.Println(msg)
}

func getGreaterOf(n1,n2 int) int{
	if n1>n2{
		return n1
	}
	return n2
}

func main(){
	Printout("Question: which is greater? 12 or 21?")
	result:=strconv.Itoa(getGreaterOf(12,21))
	Printout(result)
}
