package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//a, _ := json.Marshal(1)
	//b, _ := json.Marshal("1")
	//fmt.Println(a, string(b))

	a := struct {
		Name string
		Data []byte
	}{
		"Byte data type",
		[]byte{1,2,3},
	}

	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}





