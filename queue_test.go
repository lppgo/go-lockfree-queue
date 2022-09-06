package lfqueue

import "testing"

func TestQueueBasic(t *testing.T) {
	capacity := 10
	q := NewQueue(capacity)
	for i := 0; i < capacity; i++ {
		ok := q.Put(i)
		if !ok {
			t.Fatalf("put %d returns false", i)
		}
	}
	t.Logf("\nqueue size:%d", q.Size())
	for i := 0; i < capacity; i++ {
		elem, ok := q.Get()
		if !ok {
			t.Fatalf("get %d returns false", i)
		}
		// if elem.(int) != i {
		// 	t.Fatalf("Get wrong value")
		// }
		t.Logf("\nelem[%d]:%v\n", i, elem)
	}

	t.Log("test successed!")
}
