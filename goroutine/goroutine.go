package goroutine

import (
	"sync"
)

var wg sync.WaitGroup
var wgMap = make(map[string]interface{})
var LockMapObj = &LockMap{i: 8, lock: &sync.RWMutex{}}

// 增加一个协程
func AddGo(key string, f func(funcValue ...interface{}) interface{}, funcValue ...interface{}) {
	wg.Add(1) //添加一个计数
	 go func(funcValue ...interface{}) {
		 res := f(funcValue...)
		defer wg.Done()
		LockMapObj.Set(key, res)
	}(funcValue...)
}

//等待协程结束
func Wait() map[string]interface{} {
	wg.Wait()
	return wgMap
}

// 使用读写锁 sync.RWMutex
type LockMap struct {
	lock *sync.RWMutex
	i    int
}

func (m *LockMap) Set(k string, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	wgMap[k] = v
	return true
}

