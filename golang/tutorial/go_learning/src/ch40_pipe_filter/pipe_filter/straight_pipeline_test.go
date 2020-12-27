package pipe_filter

import "testing"

func TestStraightPipeline(t *testing.T) {
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()
	sp := NewStraightPipeline("pipeline1", spliter, converter, sum)
	ret, err := sp.Process("1,2,3")
	t.Log(ret)
	if err != nil {
		t.Fatal(err)
	}
}
