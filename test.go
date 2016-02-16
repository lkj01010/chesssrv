package main

import (
	"fmt"
//	"runtime"
//	"time"
//	"os"
//	"log"
	"encoding/json"
)

func test(c chan bool, n int) {

	x := 0
	for i := 0; i < 10000000000; i++ {
		x += i
	}

	//	fmt.Println(n, x)

	c <- true
}

type ProJsonData struct {
	UserId string
	Time   int
}

func main() {
	j1 := make(map[string]interface{})
	j1["name"] = "脚本之家"
	j1["url"] = "http://www.jb51.net/"

	{
		obj := make(map[string]interface{})
		j1["obj"] = obj
		{
			obj["p1"] = "haha"
			obj["p2"] = 1222.5
		}
	}


	//	j1["obj"] = make(map[string]interface{})
	//	j1["obj"] ["p1"] = "haha"
	//	j1["obj"] ["p2"] = 1222.5

	js1, err := json.Marshal(j1)
	if err != nil {
		panic(err)
	}

	//	jss := []byte(`{"name":"jiaoben", "url": "http:///", "obj": {"p1":"haha", "p2":123.5} }`)

	//	jd := ProJsonData{"mingzi", 5}
	//	jds, _ := json.Marshal(jd)

	println(string(js1))
	// json decode
	j2 := make(map[string]interface{})
	//	j2 := ProJsonData{}
	err = json.Unmarshal(js1, &j2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", j2)
	jj2 := (j2["obj"].(map[string]interface{}))["p2"]
	fmt.Println(jj2)
	//	fmt.Printf("%s\n", j2)
	/*
		////////////////////////////////
		// test slice
		a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		sa := a[2:7]
		fmt.Println(sa)
		sa = append(sa, 100)
		sb := sa[3:8]
		sb[0] = 99
		fmt.Println(a)  //输出：[1 2 3 4 5 99 7 100 9 0]
		fmt.Println(sa) //输出：[3 4 5 99 7 100]
		fmt.Println(sb) //输出：[99 7 100 9 0]
		//]]

		////////////////////////////////
		// test log
		file, err := os.Create("test.log")
		if err != nil {
			log.Fatalln("fail to create test.log file!")
		}
		logger := log.New(file, "[debug]", log.LstdFlags|log.Lshortfile)
		log.Println("1.Println log with log.LstdFlags ...")
		logger.Println("1.Println log with log.LstdFlags ...")
		//]]

		////////////////////////////////
		// test multi cpu
		cores := 4;
		t1 :=time.Now()
		runtime.GOMAXPROCS(cores) //设置cpu的核的数量，从而实现高并发
		c := make(chan bool)

		for i := 0; i < 1; i++ {
			go test(c, i)
		}
		//1 2.25s, 10-100-1000 1.9s, 1w 2.0s, 10w-20w 2.1s, 50w 2.57s开销， 100w就3.6了 (依次在test里面修改对应的i上限)
		for i := 0; i < 1; i++ {
			<- c
		}

		t2 :=time.Now()
		d := t2.Sub(t1)
		fmt.Println(cores, `cpu, time consume: `, d)
		//]]
	*/
}