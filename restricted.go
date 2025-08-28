package restricted

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number[T constraints.Integer | constraints.Float] struct {
	value T
	min   T
	max   T
}

func NewNumber[T constraints.Integer | constraints.Float](min, max T) *Number[T] {
	if min > max {
		min, max = max, min
	}
	return &Number[T]{
		value: min,
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
