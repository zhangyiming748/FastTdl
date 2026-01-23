package test

import (
	"fmt"
	"sync"
)

func TestQ1() {
	q1()
}

func q1() {
	var m sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i*10)
		}(i)
	}

	wg.Wait()
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}

/*
又想用协程提高效率
又必须使用线程安全锁来限制map并发
这不是很矛盾吗
*/
