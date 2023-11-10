package timer

type optionLogger struct {
	logger Logger
}

func WithLogger(logger Logger) Option {
	return &optionLogger{
		logger: logger,
	}
}

func (ol *optionLogger) apply(timer *Timer) {
	timer.logger = ol.logger
}
