package awmap

import (
	"encoding/json"
	"sync"
)

type MapInterface struct {
	StringMap map[string]interface{}
	lock      *sync.RWMutex
}

func newMapInterface(value map[string]interface{}) *MapInterface {
	m :=  &MapInterface{}
	if value == nil {
		m.StringMap = make(map[string]interface{}, 0)
	}
	m.StringMap = value
	return m
}

//线程安全添加数据
func (m *MapInterface) AddMap(k string, v interface{}) *MapInterface {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.StringMap[k] = v
	return m
}

type InArrayInterface struct {
	B bool
	Key string
	Value interface{}
}

func (m *MapInterface) InArrayValue(value int) *InArrayInterface {
	res := &InArrayInterface{B:false}
	for k,v := range m.StringMap {
		if v == value {
			res.B = true
			res.Key = k
			break
		}
	}
	return res
}

func (m *MapInterface) InArrayKey(key string) *InArrayInterface {
	res := &InArrayInterface{B:false}
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
func(i *InArrayInterface) IsOk () bool {
	return i.B
}

func (m *MapInterface) GetValue() []interface{} {
	var res []interface{}
	for _, v := range m.StringMap {
		res = append(res, v)
	}
	return res
}

func (m *MapInterface) GetKey() []string {
	var res []string
	for k, _ := range m.StringMap {
		res = append(res, k)
	}
	return res
}

func (m *MapInterface) DelByKey(key string) *MapInterface {
	delete(m.StringMap, key)
	return m
}

func (m *MapInterface) Clear() *MapInterface {
	m.StringMap = make(map[string]interface{}, 0)
	return m
}

func (m *MapInterface) ToJson() string {
	B, _ := json.Marshal(m.StringMap)
	return string(B)
}

func (m *MapInterface) Merge(list map[string]interface{}) *MapInterface {
	for k,v := range list {
		m.StringMap[k] = v
	}
	return m
}