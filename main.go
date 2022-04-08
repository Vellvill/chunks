package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"somestuff/chunk"
)

type in struct {
	Massive []int `json:"massive"`
}

func main() {
	var in in
	js, err := ioutil.ReadFile("massive.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(js, &in)
	if err != nil {
		panic(err)
	}

	out := chunk.Out{
		Data: make(map[string][]int),
	}

	out.CreateChunk(in.Massive)

	jsonn := out.Marshal()

	fmt.Println(string(jsonn))
}
