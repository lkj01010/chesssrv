package main

import (
	"fmt"
	"runtime"
	"time"
)

func test(c chan bool, n int) {

	x := 0
	for i := 0; i < 10000000000; i++ {
		x += i
	}

//	fmt.Println(n, x)

	c <- true
}

func main() {
	cores := 4;
	t1 :=time.Now()
	runtime.GOMAXPROCS(cores) //设置cpu的核的数量，从而实现高并发
	c := make(chan bool)

	for i := 0; i < 1; i++ {
		go test(c, i)
	}
	//1 2.25s, 10-100-1000 1.9s, 1w 2.0s, 10w-20w 2.1s, 50w 2.57s开销， 100w就3.6了 (依次在test里面修改对应的i上限)
	//goroutine 1个的情况下cores=4消耗2.25s，cores=1消耗7s， 叼啊！！！
	for i := 0; i < 1; i++ {
		<- c
	}

	t2 :=time.Now()
	d := t2.Sub(t1)
	fmt.Println(cores, `cpu, time consume: `, d)
}