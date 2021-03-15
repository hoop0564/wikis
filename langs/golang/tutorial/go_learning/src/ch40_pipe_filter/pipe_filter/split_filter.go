package pipe_filter

import (
	"github.com/pkg/errors"
	"strings"
)

var SplitFilerWrongFormatError = errors.New("input data should be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter: delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string) // 检查数据格式/类型，是否可以处理
	if !ok {
		return nil, SplitFilerWrongFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
