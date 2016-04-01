package main
import (
	"fmt"
	"os"
	"github.com/golang/protobuf/proto"
	"chess/com"
	"io"
)
func main() {
	// write
	msg := &com.UserInfo{
		Id:  proto.String("lkj"),
		Nickname: proto.String("mid"),
	} //msg init

	path := string("./prototestlog.txt")
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}

	defer f.Close()
	buffer, err := proto.Marshal(msg) //SerializeToOstream
	fmt.Println("marshl to=", buffer)
	f.Write(buffer)

	// read
	path = string("./prototestlog.txt")
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}

	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}
	buffer = make([]byte, fi.Size())
	_, err = io.ReadFull(file, buffer) //read all content
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}
	msg = &com.UserInfo{}
	err = proto.Unmarshal(buffer, msg) //unSerialize
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}
	fmt.Printf("read: %+v", msg)
}