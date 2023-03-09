package overview_test

import (
	"container/heap"
	"container/list"
	"testing"
)

func TestList(t *testing.T) {
	var intList list.List
	intList.PushBack(1)
	intList.PushBack(2)
	intList.PushBack(3)

	for e := intList.Front(); e != nil; e = e.Next() {
		t.Log(e.Value)
	}
}

func TestTuple(t *testing.T) {
	powerSeries := func(a int) (int, int) {
		return a * a, a * a * a
	}

	square, cube := powerSeries(3)
	if square != 9 && cube != 27 {
		t.Error("Expected 9 and 27, got ", square, cube)
	}
}

type IntegerHeap []int

func (iheap IntegerHeap) Len() int {
	return len(iheap)
}

func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}

func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

func (iheap *IntegerHeap) Push(x interface{}) {
	*iheap = append(*iheap, x.(int))
}

func (iheap *IntegerHeap) Pop() interface{} {
	old := *iheap
	n := len(old)
	x := old[n-1]
	*iheap = old[0 : n-1]
	return x
}

func TestHeaps(t *testing.T) {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	t.Logf("minimum %v", (*intHeap)[0])

	for intHeap.Len() > 0 {
		t.Log(heap.Pop(intHeap))
	}
}

func TestArray(t *testing.T) {
	var arr = [6]int{1, 2, 3, 4, 5, 6}

	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	for i, v := range arr {
		t.Log(i, v)
	}
}

func TestSlice(t *testing.T) {
	var slice = []int{1, 2, 3, 4, 5, 6}

	slice = append(slice, 7)

	t.Log("Capacity", cap(slice))
	t.Log("Length", len(slice))

	TwiceFunc := func(slice []int) {
		for i, v := range slice {
			slice[i] = v * 2
		}
	}

	TwiceFunc(slice)
	t.Log(slice)
}

func Test2DSlice(t *testing.T) {
	var slice = [][]int{
		{1, 2, 3},
		{4, 5, 6, 7, 8, 9},
		{7, 8, 9, 10, 11, 12, 13},
	}

	t.Log(slice)
}

func TestMap(t *testing.T) {
	var lanuages = map[int]string{
		1: "Go",
		2: "Java",
		3: "C++",
	}

	lanuages[4] = "Python"

	for k, v := range lanuages {
		t.Log(k, v)
	}

	delete(lanuages, 4)
	lang, exist := lanuages[3]
	if exist {
		t.Log(3, lang)
	}
}
