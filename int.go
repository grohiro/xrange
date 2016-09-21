package xrange

import (
	"strconv"
	"strings"
)

type IntRange struct {
	Start *int
	End   *int
}

func NewIntRange(from, to int) *IntRange {
	return &IntRange{
		Start: &from,
		End:   &to,
	}
}

func NewIntRangeStart(n int) *IntRange {
	return &IntRange{
		Start: &n,
		End:   nil,
	}
}

func NewIntRangeEnd(n int) *IntRange {
	return &IntRange{
		Start: nil,
		End:   &n,
	}
}

// Creates a new IntRange instance from string (format: "start,end")
func NewIntRangeFromString(str string) *IntRange {
	intRange := new(IntRange)
	if strings.Count(str, ",") == 1 {
		values := strings.Split(str, ",")
		if len(values[0]) > 0 {
			if start, err := strconv.Atoi(values[0]); err == nil {
				intRange.Start = &start
			}
		}
		if len(values[1]) > 0 {
			if end, err := strconv.Atoi(values[1]); err == nil {
				intRange.End = &end
			}
		}
	}

	return intRange
}
