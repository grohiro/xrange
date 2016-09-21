package xrange

import (
	xrange "github.com/grohiro/xrange"
	"testing"
)

func TestNewIntRange(t *testing.T) {
	r1 := xrange.NewIntRange(1, 100)
	if *r1.Start != 1 {
		t.Errorf("Expected: 1, Got: %d", *r1.Start)
	}
	if *r1.End != 100 {
		t.Errorf("Expected: 100, Got: %d", *r1.End)
	}

	r2 := xrange.NewIntRangeStart(200)
	if *r2.Start != 200 {
		t.Errorf("Expected: 200, Got: %d", r2.Start)
	}
	if r2.End != nil {
		t.Errorf("Expected: nil, Got: %p", r2.End)
	}

	r3 := xrange.NewIntRangeEnd(300)
	if r3.Start != nil {
		t.Errorf("Expected: nil, Got: %d", *r3.Start)
	}
	if *r3.End != 300 {
		t.Errorf("Expected: 300, Got: %d", *r3.End)
	}

	r4 := xrange.NewIntRangeFromString("1,100")
	if *r4.Start != 1 {
		t.Errorf("Expected: 1, Got: %d", *r4.Start)
	}
	if *r4.End != 100 {
		t.Errorf("Expected: 100, Got: %d", *r4.End)
	}

	r5 := xrange.NewIntRangeFromString("200,")
	if *r5.Start != 200 {
		t.Errorf("Expected: 200, Got: %d", r5.Start)
	}
	if r5.End != nil {
		t.Errorf("Expected: nil, Got: %p", r5.End)
	}

	r6 := xrange.NewIntRangeFromString(",300")
	if r6.Start != nil {
		t.Errorf("Expected: nil, Got: %d", *r6.Start)
	}
	if *r6.End != 300 {
		t.Errorf("Expected: 300, Got: %d", *r6.End)
	}
}
