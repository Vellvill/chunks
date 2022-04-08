package chunk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

type in struct {
	Massive []int `json:"massive"`
}

func TestOut_CreateChunk(t *testing.T) {
	t.Run("testing_test1.json", func(t *testing.T) {
		var in in
		js, err := ioutil.ReadFile("testjson/test1.json")
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(js, &in)
		if err != nil {
			panic(err)
		}
		out := Out{
			Data: make(map[string][]int),
		}

		out.CreateChunk(in.Massive)

		expected := out.Unpack()

		fmt.Println(out.Data)

		Equal(t, in.Massive, expected)

	})
	t.Run("testing_test2.json", func(t *testing.T) {
		var in in
		js, err := ioutil.ReadFile("testjson/test2.json")
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(js, &in)
		if err != nil {
			panic(err)
		}
		out := Out{
			Data: make(map[string][]int),
		}

		out.CreateChunk(in.Massive)

		expected := out.Unpack()

		fmt.Println(out.Data)

		Equal(t, in.Massive, expected)

	})
	t.Run("testing_test3.json", func(t *testing.T) {
		var in in
		js, err := ioutil.ReadFile("testjson/test3.json")
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(js, &in)
		if err != nil {
			panic(err)
		}
		out := Out{
			Data: make(map[string][]int),
		}

		out.CreateChunk(in.Massive)

		expected := out.Unpack()

		fmt.Println(out.Data)

		Equal(t, in.Massive, expected)

	})
}

func Equal(t *testing.T, expected, result interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("should be %v instead of %v", expected, result)
	}
}
