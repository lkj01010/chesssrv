package main
import (
	"time"
	"fmt"
)

func main1() {
	endless := make(chan int)

	var callback = func() {
		fmt.Print(time.Now(), "\n")
	}

	//	time.AfterFunc(3 * time.Second, callback)

	for {
		//		time.Sleep(time.Second * 1)

		<-time.After(1 * time.Second)
		callback()
	}

	<-endless
}

func main() {
	go func() {
		t := time.NewTimer(10000)
		for i := 0; i < 3; i++ {
			t.Reset(1 * time.Second)
			<-t.C
			fmt.Print(time.Now(), "\n")
		}
	}()

	time.Sleep(60 * time.Second)
}
