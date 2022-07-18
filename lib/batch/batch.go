package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {

	res = make([]user, n)
	ch := make(chan struct{}, pool)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := int64(0); i < n; i++ {
		wg.Add(1)
		mutex.Lock()
		go func(index int64) {
			ch <- struct{}{}
			//ch <- getOne(index)
			res[index] = getOne(index)
			//	res[index] = <-ch
			<-ch
			wg.Done()
		}(i)
		mutex.Unlock()
	}

	wg.Wait()
	return res
}
