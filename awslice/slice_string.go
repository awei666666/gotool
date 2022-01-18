package awslice

import (
	"strconv"
)

type SliceString struct {
	SliceString     []string
	rems            []string
	duplicateStatus bool
}

func NewSliceString(s []string) *SliceString {
	var o = &SliceString{}
	o.SliceString = s
	return o
}

func (s *SliceString) InArrayString(i string, arr ...string) bool {
	var brr []string
	brr = s.SliceString
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

func (s *SliceString) MergeString(m []string) *SliceString {
	s.SliceString = append(s.SliceString, m...)
	return s
}

func (s *SliceString) GetValue() []string {
	return s.SliceString
}

func (s *SliceString) ToIntSlice() []int {
	var res = make([]int, 0)
	for _, v := range s.SliceString {
		r, _ := strconv.Atoi(v)
		res = append(res, r)
	}
	return res
}

func (s *SliceString) SetDuplicate() *SliceString {
	s.duplicateStatus = true
	return s
}
func (s *SliceString) SetRems(rems ...string) *SliceString {
	s.rems = append(s.rems, rems...)
	return s
}

// 去重
func (s *SliceString) Rem() *SliceString {
	if s.duplicateStatus == false && len(s.rems) == 0 {
		return s
	}
	var res = make([]string, 0)
	var resMap = make(map[string]int, 0)
	for _, v := range s.SliceString {
		if s.duplicateStatus == true {
			if _, ok := resMap[v]; !ok && !s.InArrayString(v, s.rems...) {
				res = append(res, v)
				resMap[v] = 1
			}
		}else{
			if !s.InArrayString(v, s.rems...) {
				res = append(res, v)
				resMap[v] = 1
			}
		}
	}
	s.SliceString = res
	return s
}

func (s *SliceString) Search(i string) (int, bool){
	for k,v := range s.SliceString {
		if v == i {
			return k, true
		}
	}
	return 0, false
}