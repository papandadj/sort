package sort

import (
	"sync"
	"time"
)

//Sleep Sort .
// For every element x in an array, start a new program that:
// 1.Sleeps for x seconds.
// 2.Prints out x.
// The clock starts on all the elements at the same time. It works for any array that has non-negative integers.
func Sleep(arr ...int) (orders []int) {
	var wg sync.WaitGroup
	ch := make(chan int, len(arr))
	for _, item := range arr {
		wg.Add(1)
		go func(sleepTime int) {
			defer wg.Done()
			time.Sleep(time.Duration(int64(sleepTime) * time.Hour.Milliseconds()))
			ch <- sleepTime
		}(item)
	}
	wg.Wait()
	close(ch)
	for item := range ch {
		orders = append(orders, item)
	}
	return
}
