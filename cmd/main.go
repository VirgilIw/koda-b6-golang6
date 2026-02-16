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

	for data := range ch {
		mu.Lock()
		fmt.Printf("yang memesan:%s \n", data.Name)
		mu.Unlock()
		for i := data.Wait; i > 0; i-- {
			mu.Lock()
			fmt.Printf("Silahkan tunggu %d detik...\n", i)
			mu.Unlock()
			time.Sleep(1000 * time.Millisecond)
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
	go internal.Queque(ch)
	//
	wg.Add(1)
	go processQueque(ch, &wg, &mu)
	//
	wg.Wait()
}
