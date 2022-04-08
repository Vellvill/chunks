package chunk

import (
	"math"
)

type Out struct {
	Chunks [][]int `json:"chunks"`
}

func (o *Out) CreateChunk(array []int) [][]int {
	if len(array) <= 50 {
		ch := make([][]int, 0)
		ch = append(ch, array)
		return ch
	}
	ch := make([][]int, 0)
	n := math.Ceil(float64(len(array)) / 50)
	var op float64
	for op < n {
		if op == 0 {
			ch = append(ch, insert(array[int(op):int(op+50)]))
		} else if len(array[int(op)*50:]) < 50 {
			ch = append(ch, insert(array[int(op)*50:]))
		} else {
			ch = append(ch, insert(array[int(op)*50:int(op)*50+50]))
		}
		op++
	}
	return ch
}

func insert(a []int) []int {
	res := make([]int, 0)
	for _, v := range a {
		res = append(res, v)
	}
	return res
}
