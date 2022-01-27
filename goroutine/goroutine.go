package goroutine

import (
	"errors"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var wgMap = make(map[string]interface{})
var lockMapObj = &lockMap{i: 8, lock: &sync.RWMutex{}}
var wgError = make(map[string]Err)

type Err struct {
	Msg  string
	Code int // 0程序抛出的异常  1程序运行的时候抛出的异常，比如除以0，数据转换格式错误等
}

// 增加一个协程
func AddGo(key string, f func(funcValue ...interface{}) interface{}, funcValue ...interface{}) {
	wg.Add(1) //添加一个计数
	go func(funcValue ...interface{}) {
		defer func() {
			// 发生宕机时，获取panic传递的上下文并打印
			err := recover()
			if err != nil {
				switch err.(type) {
				case runtime.Error: // 运行时错误
					lockMapObj.lock.Lock()
					wgError[key] = Err{Msg: err.(runtime.Error).Error(), Code: 1}
					lockMapObj.lock.Unlock()
				default: // 非运行时错误 程序错误
					lockMapObj.lock.Lock()
					wgError[key] = Err{Msg: err.(string), Code: 0}
					lockMapObj.lock.Unlock()
				}
				lockMapObj.set(key, nil)
				defer wg.Done()
			}
		}()
		res := f(funcValue...)
		lockMapObj.set(key, res)
		defer wg.Done()
	}(funcValue...)
}

func GetValueByKey(key string) (interface{}, error) {
	value, ok := wgMap[key]
	if !ok {
		return nil, errors.New("没有找到相关协程数据")
	}
	if value == nil {
		if err, ok := wgError[key]; ok {
			var msg string
			if err.Code == 1 {
				msg = "运行时的问题:"
			} else {
				msg = "程序错误:"
			}
			return nil, errors.New(msg + err.Msg)
		}
		return nil, nil
	}
	return value, nil
}

//等待协程结束
func Wait() map[string]interface{} {
	wg.Wait()
	return wgMap
}

func GetErrMap() map[string]Err {
	return wgError
}

// 使用读写锁 sync.RWMutex
type lockMap struct {
	lock *sync.RWMutex
	i    int
}

func (m *lockMap) set(k string, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	wgMap[k] = v
	return true
}
