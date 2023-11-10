package timer

type optionMethod struct {
	method string
}

func WithMethod(method string) Option {
	return &optionMethod{
		method: method,
	}
}

func (om *optionMethod) apply(timer *Timer) {
	timer.method = om.method
}
