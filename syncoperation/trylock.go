package syncoperation

import (
	"fmt"
	"sync"
)

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <- l.c:
		lockResult = true
	default:
	}
	return lockResult
}

func (l Lock) Unlock() {
	l.c <- struct{}{}
}

func Counter() {
	var counter int
	var l = NewLock()
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				fmt.Println("lock failed")
				return
			}
			counter++
			l.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}