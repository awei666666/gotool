package awmap

import (
	"encoding/json"
	"sync"
)

type MapIntInterface struct {
	StringMap map[int]interface{}
	lock      *sync.RWMutex
}

func NewMapIntInterface(value map[int]interface{}) *MapIntInterface {
	m :=  &MapIntInterface{}
	if value == nil {
		m.StringMap = make(map[int]interface{}, 0)
	}
	m.StringMap = value
	return m
}

//线程安全添加数据
func (m *MapIntInterface) AddMap(k int, v interface{}) *MapIntInterface {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.StringMap[k] = v
	return m
}

type InArrayIntInterface struct {
	B bool
	Key int
	Value interface{}
}

func (m *MapIntInterface) InArrayValue(value interface{}) *InArrayIntInterface {
	res := &InArrayIntInterface{B:false}
	for k,v := range m.StringMap {
		if v == value {
			res.B = true
			res.Key = k
			break
		}
	}
	return res
}

func (m *MapIntInterface) InArrayKey(key int) *InArrayIntInterface {
	res := &InArrayIntInterface{B:false}
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
func(i *InArrayIntInterface) IsOk () bool {
	return i.B
}

func (m *MapIntInterface) GetValue() []interface{} {
	var res []interface{}
	for _, v := range m.StringMap {
		res = append(res, v)
	}
	return res
}

func (m *MapIntInterface) GetKey() []int {
	var res []int
	for k, _ := range m.StringMap {
		res = append(res, k)
	}
	return res
}

func (m *MapIntInterface) DelByKey(key int) *MapIntInterface {
	delete(m.StringMap, key)
	return m
}

func (m *MapIntInterface) Clear() *MapIntInterface {
	m.StringMap = make(map[int]interface{}, 0)
	return m
}

func (m *MapIntInterface) ToJson() string {
	B, _ := json.Marshal(m.StringMap)
	return string(B)
}

func (m *MapIntInterface) Merge(list map[int]interface{}) *MapIntInterface {
	for k,v := range list {
		m.StringMap[k] = v
	}
	return m
}