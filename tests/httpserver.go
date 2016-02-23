package main
import (
	"fmt"
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir("."))
//	var port string
//	fmt.Scanf("%s", &port)
	port := "13777"
	fmt.Println("serving charles http on port:",port)
	http.ListenAndServe(":" + port, h)
}
