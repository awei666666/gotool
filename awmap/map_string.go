package awmap

import (
	"encoding/json"
	"sync"
)

type MapString struct {
	StringMap map[string]string
	lock      *sync.RWMutex
}

func newMapString(value map[string]string) *MapString {
	m :=  &MapString{}
	if value == nil {
		m.StringMap = make(map[string]string, 0)
	}
	m.StringMap = value
	return m
}

//线程安全添加数据
func (m *MapString) AddMap(k string, v string) *MapString {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.StringMap[k] = v
	return m
}

type InArrayString struct {
	B bool
	Key string
	Value string
}

func (m *MapString) InArrayValue(value string) *InArrayString {
	res := &InArrayString{B:false}
	for k,v := range m.StringMap {
		if v == value {
			res.B = true
			res.Key = k
			break
		}
	}
	return res
}

func (m *MapString) InArrayKey(key string) *InArrayString {
	res := &InArrayString{B:false}
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
func(i *InArrayString) IsOk () bool {
	return i.B
}

func (m *MapString) GetValue() []string {
	var res []string
	for _, v := range m.StringMap {
		res = append(res, v)
	}
	return res
}

func (m *MapString) GetKey() []string {
	var res []string
	for k, _ := range m.StringMap {
		res = append(res, k)
	}
	return res
}

func (m *MapString) DelByKey(key string) *MapString {
	delete(m.StringMap, key)
	return m
}

func (m *MapString) Clear() *MapString {
	m.StringMap = make(map[string]string, 0)
	return m
}

func (m *MapString) ToJson() string {
	B, _ := json.Marshal(m.StringMap)
	return string(B)
}

func (m *MapString) Merge(list map[string]string) *MapString {
	for k,v := range list {
		m.StringMap[k] = v
	}
	return m
}