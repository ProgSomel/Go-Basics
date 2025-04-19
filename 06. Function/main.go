package main

import "fmt"

func call() func(int, int){
	return add;
}

func add(x int, y int){
	z:= x + y;
	fmt.Println(z);
}
func main(){
	sum:= call(); // add function
	sum(4, 5) // 9
}