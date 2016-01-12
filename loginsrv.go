package main

import (
	"log"
	"os"
	"fmt"
)

func main() {
	file, err := os.Create("test.log")
	if err != nil {
	log.Fatalln("fail to create test.log file!")
	}
	logger := log.New(file, "[debug]", log.LstdFlags|log.Lshortfile)
	log.Println("1.Println log with log.LstdFlags ...")
	logger.Println("1.Println log with log.LstdFlags ...")

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	sa := a[2:7]
	fmt.Println(sa)
	sa = append(sa, 100)
	sb := sa[3:8]
	sb[0] = 99
	fmt.Println(a)  //输出：[1 2 3 4 5 99 7 100 9 0]
	fmt.Println(sa) //输出：[3 4 5 99 7 100]
	fmt.Println(sb) //输出：[99 7 100 9 0]


	//	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	//		s := http.Server{Handler: websocket.Handler(wsHandler)}
	//		s.ServeHTTP(w, req)
	//	})
	//
	//	err := http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil)
	//	if err != nil {
	//		panic("Error: " + err.Error())
	//	}


}