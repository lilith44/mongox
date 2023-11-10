package timer

type optionCollection struct {
	collection string
}

func WithCollection(collection string) Option {
	return &optionCollection{
		collection: collection,
	}
}

func (oc *optionCollection) apply(timer *Timer) {
	timer.collection = oc.collection
}
