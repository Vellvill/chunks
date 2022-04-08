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

	var out chunk.Out

	out.Chunks = out.CreateChunk(in.Massive)

	res, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}
