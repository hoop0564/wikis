package pipe_filter

import (
	"github.com/pkg/errors"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("input data should be []string")

type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (sf *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string) // 检查数据格式/类型，是否可以处理
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}
	var ret []int
	for _, part := range parts {
		s, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}
	return ret, nil
}
