package awmap

import (
"encoding/json"
"sync"
)

type MapInt struct {
	StringMap map[string]int
	lock      *sync.RWMutex
}

func newMapInt(value map[string]int) *MapInt {
	m :=  &MapInt{}
	if value == nil {
		m.StringMap = make(map[string]int, 0)
	}
	m.StringMap = value
	return m
}

//线程安全添加数据
func (m *MapInt) AddMap(k string, v int) *MapInt {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.StringMap[k] = v
	return m
}

type InArrayInt struct {
	B bool
	Key string
	Value int
}

func (m *MapInt) InArrayValue(value int) *InArrayInt {
	res := &InArrayInt{B:false}
	for k,v := range m.StringMap {
		if v == value {
			res.B = true
			res.Key = k
			break
		}
	}
	return res
}

func (m *MapInt) InArrayKey(key string) *InArrayInt {
	res := &InArrayInt{B:false}
	for k,v := range m.StringMap {
		if k == key {
			res.B = true
			res.Value = v
			break
		}
	}
	return res
}

// 返回inArray的结果，如果是true 则证明需要判断的值，存在在map中
func(i *InArrayInt) IsOk () bool {
	return i.B
}

func (m *MapInt) GetValue() []int {
	var res []int
	for _, v := range m.StringMap {
		res = append(res, v)
	}
	return res
}

func (m *MapInt) GetKey() []string {
	var res []string
	for k, _ := range m.StringMap {
		res = append(res, k)
	}
	return res
}

func (m *MapInt) DelByKey(key string) *MapInt {
	delete(m.StringMap, key)
	return m
}

func (m *MapInt) Clear() *MapInt {
	m.StringMap = make(map[string]int, 0)
	return m
}

func (m *MapInt) ToJson() string {
	B, _ := json.Marshal(m.StringMap)
	return string(B)
}

func (m *MapInt) Merge(list map[string]int) *MapInt {
	for k,v := range list {
		m.StringMap[k] = v
	}
	return m
}