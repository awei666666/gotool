package awmap

import (
	"encoding/json"
	"sync"
)

type MapIntInt struct {
	StringMap map[int]int
	lock      *sync.RWMutex
}

func NewMapIntInt(value map[int]int) *MapIntInt {
	m :=  &MapIntInt{}
	if value == nil {
		m.StringMap = make(map[int]int, 0)
	}
	m.StringMap = value
	return m
}

//线程安全添加数据
func (m *MapIntInt) AddMap(k int, v int) *MapIntInt {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.StringMap[k] = v
	return m
}

type InArrayIntInt struct {
	B bool
	Key int
	Value int
}

func (m *MapIntInt) InArrayValue(value int) *InArrayIntInt {
	res := &InArrayIntInt{B:false}
	for k,v := range m.StringMap {
		if v == value {
			res.B = true
			res.Key = k
			break
		}
	}
	return res
}

func (m *MapIntInt) InArrayKey(key int) *InArrayIntInt {
	res := &InArrayIntInt{B:false}
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
func(i *InArrayIntInt) IsOk () bool {
	return i.B
}

func (m *MapIntInt) GetValue() []int {
	var res []int
	for _, v := range m.StringMap {
		res = append(res, v)
	}
	return res
}

func (m *MapIntInt) GetKey() []int {
	var res []int
	for k, _ := range m.StringMap {
		res = append(res, k)
	}
	return res
}

func (m *MapIntInt) DelByKey(key int) *MapIntInt {
	delete(m.StringMap, key)
	return m
}

func (m *MapIntInt) Clear() *MapIntInt {
	m.StringMap = make(map[int]int, 0)
	return m
}

func (m *MapIntInt) ToJson() string {
	B, _ := json.Marshal(m.StringMap)
	return string(B)
}

func (m *MapIntInt) Merge(list map[int]int) *MapIntInt {
	for k,v := range list {
		m.StringMap[k] = v
	}
	return m
}