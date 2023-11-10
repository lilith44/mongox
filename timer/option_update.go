package timer

type optionUpdate struct {
	update any
}

func WithUpdate(update any) Option {
	return &optionUpdate{
		update: update,
	}
}

func (ou *optionUpdate) apply(timer *Timer) {
	timer.update = ou.update
}
