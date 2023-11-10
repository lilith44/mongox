package timer

type optionPipeline struct {
	pipeline any
}

func WithPipeline(pipeline any) Option {
	return &optionPipeline{
		pipeline: pipeline,
	}
}

func (om *optionPipeline) apply(timer *Timer) {
	timer.pipeline = om.pipeline
}
