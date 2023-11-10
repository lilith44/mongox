package timer

import "go.mongodb.org/mongo-driver/mongo"

type optionModel struct {
	models []mongo.WriteModel
}

func WithModel(models []mongo.WriteModel) Option {
	return &optionModel{
		models: models,
	}
}

func (om *optionModel) apply(timer *Timer) {
	timer.models = om.models
}
