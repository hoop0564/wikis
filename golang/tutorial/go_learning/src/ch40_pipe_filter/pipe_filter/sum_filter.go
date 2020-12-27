package pipe_filter

import (
	"github.com/pkg/errors"
)

var SumFilterWrongFormatError = errors.New("input data should be []int")

type SumFilter struct {
}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
	elems, ok := data.([]int) // 检查数据格式/类型，是否可以处理
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	ret := 0
	for _, elem := range elems {
		ret += elem
	}
	return ret, nil
}
