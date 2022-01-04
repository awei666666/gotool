package awmap

import (
	"encoding/json"
	"sync"
)

type MapIntString struct {
	StringMap map[int]string
	lock      *sync.RWMutex
}

func NewMapIntString(value map[int]string) *MapIntString {
	m :=  &MapIntString{}
	if value == nil {
		m.StringMap = make(map[int]string, 0)
	}
	m.StringMap = value
	return m
}

//线程安全添加数据
func (m *MapIntString) AddMap(k int, v string) *MapIntString {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.StringMap[k] = v
	return m
}

type InArrayIntString struct {
	B bool
	Key int
	Value string
}

func (m *MapIntString) InArrayValue(value string) *InArrayIntString {
	res := &InArrayIntString{B:false}
	for k,v := range m.StringMap {
		if v == value {
			res.B = true
			res.Key = k
			break
		}
	}
	return res
}

func (m *MapIntString) InArrayKey(key int) *InArrayIntString {
	res := &InArrayIntString{B:false}
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
func(i *InArrayIntString) IsOk () bool {
	return i.B
}

func (m *MapIntString) GetValue() []string {
	var res []string
	for _, v := range m.StringMap {
		res = append(res, v)
	}
	return res
}

func (m *MapIntString) GetKey() []int {
	var res []int
	for k, _ := range m.StringMap {
		res = append(res, k)
	}
	return res
}

func (m *MapIntString) DelByKey(key int) *MapIntString {
	delete(m.StringMap, key)
	return m
}

func (m *MapIntString) Clear() *MapIntString {
	m.StringMap = make(map[int]string, 0)
	return m
}

func (m *MapIntString) ToJson() string {
	B, _ := json.Marshal(m.StringMap)
	return string(B)
}

func (m *MapIntString) Merge(list map[int]string) *MapIntString {
	for k,v := range list {
		m.StringMap[k] = v
	}
	return m
}