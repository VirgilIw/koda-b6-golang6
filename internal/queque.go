package internal

import "sync"

func Queque(ch chan Data, wg *sync.WaitGroup) {
	defer wg.Done()
	data := []Data{
		{Name: "vir", Wait: 4},
		{Name: "ver", Wait: 3},
		{Name: "vur", Wait: 6},
		{Name: "vor", Wait: 2},
		{Name: "vmr", Wait: 1},
	}

	for _, x := range data {
		ch <- x
	}
}
