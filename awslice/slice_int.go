package awslice

import (
	"strconv"
	"sync"
)

type SliceInt struct {
	SliceInt        []int
	lock            *sync.RWMutex
	rems            []int
	duplicateStatus bool
}

func NewSliceInt(s []int) *SliceInt {
	var o = &SliceInt{}
	o.SliceInt = s
	return o
}

func (s *SliceInt) InArrayInt(i int, arr ...int) bool {
	var brr []int
	brr = s.SliceInt
	if len(arr) > 0 {
		brr = arr
	}
	for _, v := range brr {
		if i == v {
			return true
		}
	}
	return false
}
func (s *SliceInt) SetDuplicate() *SliceInt {
	s.duplicateStatus = true
	return s
}
func (s *SliceInt) SetRems(rems ...int) *SliceInt {
	s.rems = append(s.rems, rems...)
	return s
}

// 去重
func (s *SliceInt) Rem() *SliceInt {
	if s.duplicateStatus == false && len(s.rems) == 0 {
		return s
	}
	var res = make([]int, 0)
	var resMap = make(map[int]int, 0)
	for _, v := range s.SliceInt {
		if s.duplicateStatus == true {
			if _, ok := resMap[v]; !ok && !s.InArrayInt(v, s.rems...) {
				res = append(res, v)
				resMap[v] = 1
			}
		} else {
			if !s.InArrayInt(v, s.rems...) {
				res = append(res, v)
				resMap[v] = 1
			}
		}
	}
	s.SliceInt = res
	return s
}

// 合并 并集
func (s *SliceInt) MergeInt(m []int) *SliceInt {
	s.SliceInt = append(s.SliceInt, m...)
	return s
}

func (s *SliceInt) GetValue() []int {
	return s.SliceInt
}

func (s *SliceInt) ToStringSlice() []string {
	var res = make([]string, 0)
	for _, v := range s.SliceInt {
		res = append(res, strconv.Itoa(v))
	}
	return res
}

// 排序  正序  倒序
func (s *SliceInt) SortDesc() *SliceInt {
	s.SliceInt = s.sort(s.SliceInt, true)
	return s
}

func (s *SliceInt) SortAsc() *SliceInt {
	s.SliceInt = s.sort(s.SliceInt, false)
	return s
}

func (s *SliceInt) sort(parInt []int, b bool) []int {
	var count = len(parInt)
	if count == 1 {
		return parInt
	}
	mid := count / 2
	leftInt := parInt[0:mid]
	rightInt := parInt[mid:]
	leftInt = s.sort(leftInt, b)
	rightInt = s.sort(rightInt, b)
	return s.merge(leftInt, rightInt, b)
}

func (s *SliceInt) merge(le, ri []int, b bool) []int {
	var res = make([]int, 0)
	for {
		if len(le) == 0 || len(ri) == 0 {
			break
		}
		if b == true {
			if le[0] > ri[0] {
				res = append(res, le[0])
				le = le[1:]
			} else {
				res = append(res, ri[0])
				ri = ri[1:]
			}
		} else {
			if le[0] < ri[0] {
				res = append(res, le[0])
				le = le[1:]
			} else {
				res = append(res, ri[0])
				ri = ri[1:]
			}
		}
	}
	if len(le) > 0 {
		res = append(res, le...)
	}
	if len(ri) > 0 {
		res = append(res, ri...)
	}
	return res
}

func (s *SliceInt) Search(i int) (int, bool) {
	for k, v := range s.SliceInt {
		if v == i {
			return k, true
		}
	}
	return 0, false
}

// 并集
func (s *SliceInt) Intersection(data []int) []int {
	var res = make([]int, 0)
	var resMap = make(map[int]int, 0)
	for _, v := range s.SliceInt {
		resMap[v] = 1
	}
	for _, v := range data {
		if _, ok := resMap[v]; ok {
			res = append(res, v)
		}
	}
	return res
}

// 差集
func (s *SliceInt) Diff(data []int) []int {
	var res = make([]int, 0)
	var resMap = make(map[int]int, 0)
	for _, v := range s.SliceInt {
		resMap[v] = 1
	}
	for _, v := range data {
		if _, ok := resMap[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}
