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

//	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
//		s := http.Server{Handler: websocket.Handler(wsHandler)}
//		s.ServeHTTP(w, req)
//	}
//	err := http.ListenAndServe(":"+strconv.Itoa(Config.Port), nil)
//	if err != nil {
//		panic("Error: " + err.Error())
//	}
