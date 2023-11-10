package timer

import (
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Timer struct {
	startAt    time.Time
	collection string
	method     string
	bean       any
	filter     any
	update     any
	pipeline   any
	models     []mongo.WriteModel
	logger     Logger
}

type Logger interface {
	Info(args ...any)
	Infof(template string, args ...any)
	Errorf(template string, args ...any)
}

func New(options ...Option) *Timer {
	timer := &Timer{
		startAt: time.Now(),
	}

	for _, option := range options {
		option.apply(timer)
	}
	timer.start()
	return timer
}

func (t *Timer) start() {
	if t.logger == nil {
		return
	}

	var fields []string
	if t.collection != "" {
		fields = append(fields, "operate collection "+t.collection)
	}
	if t.method != "" {
		fields = append(fields, "method="+t.method)
	}
	if t.bean != nil {
		fields = append(fields, fmt.Sprintf("bean=%+v", t.bean))
	}
	if t.filter != nil {
		fields = append(fields, fmt.Sprintf("filter=%+v", t.filter))
	}
	if t.update != nil {
		fields = append(fields, fmt.Sprintf("update=%+v", t.update))
	}
	if t.models != nil {
		fields = append(fields, fmt.Sprintf("models=%+v", t.models))
	}
	if t.pipeline != nil {
		fields = append(fields, fmt.Sprintf("pipeline=%+v", t.pipeline))
	}

	t.logger.Info(strings.Join(fields, ", "))
}

func (t *Timer) End(err error) {
	if t.logger == nil {
		return
	}

	cost := time.Since(t.startAt).Milliseconds()
	if err != nil {
		t.logger.Errorf("%s failed, error=%s, cost=%dms", t.method, err, cost)
		return
	}
	t.logger.Infof("%s succeeded, cost=%dms", t.method, cost)
}
