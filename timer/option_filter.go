package timer

type optionFilter struct {
	filter any
}

func WithFilter(filter any) Option {
	return &optionFilter{
		filter: filter,
	}
}

func (of *optionFilter) apply(timer *Timer) {
	timer.filter = of.filter
}
