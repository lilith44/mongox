package timer

type optionBean struct {
	bean any
}

func WithBean(bean any) Option {
	return &optionBean{
		bean: bean,
	}
}

func (ob *optionBean) apply(timer *Timer) {
	timer.bean = ob.bean
}
