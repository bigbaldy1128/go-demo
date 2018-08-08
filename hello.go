package main

import (
	"fmt"
	"strconv"
	"time"
)

type I interface{
	show() string
}

type A struct{
	name string
	age int
}

type B struct{}

func (*B) show() string{
	return "B"
}

func (a A) show() string{
	return "A"
}

func main() {
	var a B
	fmt.Println(a.show())
}

func calc(num int,result chan int){
	result <- fab(num)
}

func test1(){
	start := time.Now().UnixNano();
	result1 := make(chan int)
	result2 := make(chan int)
	go calc(44,result1)
	go calc(43,result2)
	result := strconv.Itoa(<-result1 + <-result2)
	end := time.Now().UnixNano();
	elapseTime := strconv.FormatInt((end-start)/1e6,10)
	fmt.Printf("elapseTime:"+elapseTime+",result:"+result)
}

func test2() chan int{
	out := make(chan int)
	out <- 2
	return out
}

func fab(num int) int {
	if num < 2 {
		return 1
	}
	return fab(num-1)+fab(num-2)
}