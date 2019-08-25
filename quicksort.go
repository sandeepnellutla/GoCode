package main

import(
	"fmt"
	"math/rand"
)

func initialize() []int{
	a:= []int{35,12,2,21,1}
	return a
}

func sortem(values []int)[]int{
	if len(values)<2{
		return values
	}

	left,right:=0,len(values)-1
	pivot:=rand.Int()%len(values)
	values[pivot],values[right]=values[right],values[pivot]
	for i,_:=range values{
		if values[i]<values[right]{
			values[left],values[i]=values[i],values[left]
			left++
		}
	}

	values[left],values[right]=values[right],values[left]

	sortem(values[:left])
	sortem(values[left+1:])

	return values
}

func main(){
	values:=initialize()
	fmt.Println("\n---------- Unsorted--------\n",values)
	sortem(values)
	fmt.Println("\n---------- Sorted ----------\n",values)
}
