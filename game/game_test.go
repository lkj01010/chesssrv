package game

import (
	"testing"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func TestJson(t *testing.T) {


}
func readJsonFile(filename string) (map[string]string, error) {
	var xxx = map[string]string{}


	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}

	if err := json.Unmarshal(bytes, &xxx); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}

	return xxx, nil
}

func TestReadJsonFile(t *testing.T)  {
	filename := "./cmd.json"
	xxxMap, err := readJsonFile(filename)
	if err != nil {
		fmt.Println("readFile: ", err.Error())
		return
	}

	fmt.Println("xxxmap:", xxxMap)
}