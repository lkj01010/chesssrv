package com
import (
	"github.com/lkj01010/log"
	"path/filepath"
	"io/ioutil"
	"fmt"
	"bytes"
)

/**
 * generate external command file for client using
 */
func GenCommandOutput(path string){
	var buf bytes.Buffer
	for i := Cmd_Ag_Start; i <= Cmd_Ag_End; i ++ {
		fmt.Fprintf(&buf, "%v = %d\n", i, i)
	}
	fmt.Fprint(&buf, "\n")

	for i := Cmd_Game_Start; i <= Cmd_Game_End; i ++ {
		fmt.Fprintf(&buf, "%v = %d\n", i, i)
	}
	fmt.Fprint(&buf, "\n")

	for i := Cmd_Cow_Start; i <= Cmd_Cow_End; i ++ {
		fmt.Fprintf(&buf, "%v = %d\n", i, i)
	}

	err := ioutil.WriteFile(filepath.Join(path, "./command.txt"), buf.Bytes(), 0666)  //写入文件(字节数组)
	if err != nil {
		log.Error(err)
	}
}