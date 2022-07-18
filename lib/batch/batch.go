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
	var waitingGroup sync.WaitGroup
	for i := int64(0); i < n; i++ {
		waitingGroup.Add(1)
		ch <- struct{}{}
		go func(index int64) {
			res[index] = getOne(index)
			<-ch
			waitingGroup.Done()
		}(i)

	}
	waitingGroup.Wait()
	return res
}
