package restricted

import (
	"flag"
	"fmt"

	"golang.org/x/exp/constraints"
)

var (
	_ flag.Value = (*Number[int])(nil)
	_ flag.Value = (*String)(nil)
)

type Number[T constraints.Integer | constraints.Float] struct {
	value T
	min   T
	max   T
}

func NewNumber[T constraints.Integer | constraints.Float](value, min, max T) *Number[T] {
	if min > max {
		min, max = max, min
	}
	return &Number[T]{
		value: value,
		min:   min,
		max:   max,
	}
}

func (r *Number[T]) String() string {
	return fmt.Sprintf("%v", r.value)
}

func (r *Number[T]) Set(s string) error {
	var val T
	_, err := fmt.Sscanf(s, "%v", &val)
	if err != nil {
		return fmt.Errorf("invalid number: %s", s)
	}
	if val < r.min || val > r.max {
		return fmt.Errorf("value must be between %v and %v, got: %v", r.min, r.max, val)
	}
	r.value = val
	return nil
}

func (r *Number[T]) Min() T {
	return r.min
}

func (r *Number[T]) Max() T {
	return r.max
}

func (r *Number[T]) Get() T {
	return r.value
}

type String struct {
	value string
	min   int
	max   int
}

func NewString(value string, min, max int) *String {
	if min > max {
		min, max = max, min
	}
	return &String{
		value: value,
		min:   min,
		max:   max,
	}
}

func (r *String) String() string {
	return r.value
}

func (r *String) Set(s string) error {
	if len(s) < r.min || len(s) > r.max {
		return fmt.Errorf("length must be between %v and %v, got: %v", r.min, r.max, len(s))
	}
	r.value = s
	return nil
}

func (r *String) Min() int {
	return r.min
}

func (r *String) Max() int {
	return r.max
}

func (r *String) Get() string {
	return r.value
}
