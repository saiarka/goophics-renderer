package main

type Interval struct {
	min              float64
	max              float64
	is_min_inclusive bool
	is_max_inclusive bool
}

func NewInterval(min float64, max float64, is_min_inclusive bool, is_max_inclusive bool) Interval {
	return Interval{
		min:              min,
		max:              max,
		is_min_inclusive: is_min_inclusive,
		is_max_inclusive: is_max_inclusive,
	}
}

func (i *Interval) InInterval(t float64) (is_in_interval bool) {
	if i.is_min_inclusive {
		if t < i.min {
			return false
		}
	} else {
		if t <= i.min {
			return false
		}
	}

	if i.is_max_inclusive {
		if t > i.max {
			return false
		}
	} else {
		if t >= i.max {
			return false
		}
	}

	return true
}

func (i *Interval) SetInterval(min float64, max float64, is_min_inclusive bool, is_max_inclusive bool) {
	i.min = min
	i.max = max
	i.is_min_inclusive = is_min_inclusive
	i.is_max_inclusive = is_max_inclusive
}

func (i *Interval) Size() (size float64) {
	return i.max - i.min
}
