package implement_queue_using_stacks

import "testing"

func Test_case1_success(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	p1 := obj.Peek()
	if 1 != p1 {
		t.Errorf("%v != %v ", 1, p1)
	}
	p2 := obj.Pop()
	if 1 != p2 {
		t.Errorf("%v != %v ", 2, p2)
	}
	p3 := obj.Empty()
	if false != p3 {
		t.Errorf("%v != %v ", false, p3)
	}
}

func Test_case2_success(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	p1 := obj.Pop()
	if 1 != p1 {
		t.Errorf("%v != %v ", 1, p1)
	}
	obj.Push(5)
	p2 := obj.Pop()
	if 2 != p2 {
		t.Errorf("%v != %v ", 2, p2)
	}
	p3 := obj.Pop()
	if 3 != p3 {
		t.Errorf("%v != %v ", 3, p3)
	}
	p4 := obj.Pop()
	if 4 != p4 {
		t.Errorf("%v != %v ", 4, p4)
	}
	p5 := obj.Pop()
	if 5 != p5 {
		t.Errorf("%v != %v ", 5, p5)
	}
}
