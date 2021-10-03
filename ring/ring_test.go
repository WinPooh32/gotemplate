package ringfixed

import (
	"reflect"
	"testing"
)

func TestRing_CopyTo(t *testing.T) {
	const capacity = 9

	var ring = Make(capacity)
	var data = []DType{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	for _, v := range data {
		if ring.Full() {
			t.Logf("pop front %v", ring.PopFront())
		}
		ring.PushBack(v)
		t.Logf("push back %v", v)
	}

	t.Logf("ring internal buf: %v", ring.buf)

	var buf = make([]DType, capacity)
	n := ring.CopyTo(buf)
	t.Logf("result buf: %v", buf[:n])

	want := []DType{8, 9, 10, 11, 12, 13, 14, 15, 16}
	got := buf

	if !reflect.DeepEqual(got, want) {
		t.Logf("want: %+v, got: %+v", want, got)
		t.Fail()
		return
	}
}
