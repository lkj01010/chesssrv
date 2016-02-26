package main
import "fmt"

func main(){
	//test evaluation
	mapa := make(map[int]int)
	mapa[1] = 100
	mapa[2] = 200

	mapb := mapa	//这里只是地址复制，非值拷贝
	mapb[1] = 5000

	fmt.Printf("mapa=%+v, mapb=%+v\n", mapa, mapb)

	// test delete
	delete(mapb, 1)
	fmt.Printf("now is %+v\n", mapb)

	stringmap := make(map[string]string)
	stringmap["aaa"] = "stringaaa"
	stringmap["bbb"] = "stringbbb"
	fmt.Printf("stringmap=%v\n", stringmap)
	delete(stringmap, "aaa")
	fmt.Printf("stringmap=%v\n", stringmap)
}
