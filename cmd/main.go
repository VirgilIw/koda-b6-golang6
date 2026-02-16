package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/virgilIw/koda-b6-go6/internal"
)

// mutex itu mengunci akses ke bagian kode tertentu supaya hanya satu goroutine yang bisa masuk pada saat yang sama.

func processQueque(ch chan internal.Data, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	data := 5
	for range data {
		data := <-ch

		mu.Lock()
		fmt.Printf("yang memesan:%s \n", data.Name)
		mu.Unlock()

		for j := data.Wait; j > 0; j-- {
			mu.Lock()
			fmt.Printf("Silahkan tunggu %d detik...\n", j)
			mu.Unlock()
			time.Sleep(1 * time.Second)
		}

		mu.Lock()
		fmt.Println("Selesai:", data.Name)
		fmt.Println("------------------")
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan internal.Data)
	wg.Add(2)
	go internal.Queque(ch, &wg)
	go processQueque(ch, &wg, &mu)
	//
	wg.Wait()
}
