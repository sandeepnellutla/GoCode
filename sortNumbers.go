package main

import(
	"fmt"
)

func initializeNumbers() [5]int{
	b:= [5]int{3,5,21,23,55}
	return b
}

func Printstr(msg string){
	fmt.Printf("%v",msg)
}

func Printnum(n int){
	fmt.Printf("%d ",n)
}

func getGreaterOf(n1,n2 int)int{
	if n1>n2{
		return n1
	}
	return n2
}

func main(){
	unsorted:=initializeNumbers()
	Printstr("Sorting...")
	for k:=range unsorted{
		Printnum(k)
	}
	Printstr("\n")
}
