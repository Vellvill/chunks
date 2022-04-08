package chunk

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"sync"
)

type Out struct {
	sync.Mutex
	Data map[string][]int `json:"chunks"`
}

func (o *Out) CreateChunk(array []int) {
	wg := new(sync.WaitGroup)
	if len(array) <= 50 {
		if len(array) == 0 {
			err := fmt.Errorf("Nil len")
			if err != nil {
				log.Fatal(err)
			}
		}
		ch := make([][]int, 0)
		ch = append(ch, array)
		o.Data["0"] = array
	}
	n := math.Ceil(float64(len(array)) / 50)
	var op float64
	for op < n {
		wg.Wait()
		if op == 0 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				o.Lock()
				o.Data[strconv.Itoa(int(op))] = insert(array[int(op):int(op+50)])
				op++
				o.Unlock()
			}()
		} else if len(array[int(op)*50:]) < 50 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				o.Lock()
				o.Data[strconv.Itoa(int(op))] = insert(array[int(op)*50:])
				o.Unlock()
			}()
			break
		} else {
			wg.Add(1)
			go func() {
				defer wg.Done()
				o.Lock()
				o.Data[strconv.Itoa(int(op))] = insert(array[int(op)*50 : int(op)*50+50])
				op++
				o.Unlock()
			}()
		}
	}
	wg.Wait()
}

func (o *Out) Marshal() []byte {
	res, err := json.Marshal(o)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func insert(a []int) []int {
	res := make([]int, 0)
	for _, v := range a {
		res = append(res, v)
	}
	return res
}
