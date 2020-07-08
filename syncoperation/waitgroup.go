package syncoperation

import (
	"fmt"
	"sync"
)


func waitgroup() {
	var wg sync.WaitGroup
	var counter int
	var l sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++
			l.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
