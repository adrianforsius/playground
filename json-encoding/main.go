package main

import (
	"bytes"

	"fmt"

	"encoding/json"
)

type Table struct {
	Max string `json:"max"`
}

func main() {
	t := Table{Max: "hello"}
	table := new(bytes.Buffer)
	err := json.NewEncoder(table).Encode(t)
	if err != nil {
		panic(err)
	}
	fmt.Println("test", table.Bytes())

}
