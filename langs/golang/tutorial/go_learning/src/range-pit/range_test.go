package range_pit

import (
	"testing"
)

func TestRangeSlice(t *testing.T) {
	arr := []int{1, 2, 3}
	var newArr []*int
	for _, v := range arr {
		newArr = append(newArr, &v)
	}
	for _, v := range newArr {
		t.Log(*v)
	}
}

// 输出 3 3 3
// 因为for range在循环时，go会创建一个额外的变量去存储循环的元素，所以在每一次迭代中，该变量都会被重新赋值，由于这里使用的是指针，所以就出现上面的这种情况。我们可以用&arr[i]去替代&v

func TestModifyArray(t *testing.T) {
	a := []int{1,2,3}
	for _, v := range a {
		a = append(a, v)
	}
	t.Log(a)
}

// 输出 [1 2 3 1 2 3]
// 因为for range在编译期间，就会把a赋值给一个新的变量，所以我们遍历的其实已经不是a变量了。