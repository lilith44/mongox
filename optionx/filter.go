package optionx

type Filter struct {
	filter any
}

func WithFilter(filter any) *Filter {
	return &Filter{
		filter: filter,
	}
}

func (f *Filter) ApplyCount(options *CountOptions) {
	options.Filter = f.filter
}

func (f *Filter) ApplyFind(options *FindOptions) {
	options.Filter = f.filter
}

func (f *Filter) ApplyFindOne(options *FindOneOptions) {
	options.Filter = f.filter
}

func (f *Filter) ApplyReplaceOne(options *ReplaceOneOptions) {
	options.Filter = f.filter
}

func (f *Filter) ApplyUpdateMany(options *UpdateManyOptions) {
	options.Filter = f.filter
}

func (f *Filter) ApplyUpdateOne(options *UpdateOneOptions) {
	options.Filter = f.filter
}
