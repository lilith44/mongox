package optionx

type Unscoped struct {
	unscoped bool
}

func WithUnscoped() *Unscoped {
	return &Unscoped{
		unscoped: true,
	}
}

func (u *Unscoped) ApplyAggregate(options *AggregateOptions) {
	options.Unscoped = u.unscoped
}

func (u *Unscoped) ApplyCount(options *CountOptions) {
	options.Unscoped = u.unscoped
}

func (u *Unscoped) ApplyDeleteOne(options *DeleteOneOptions) {
	options.Unscoped = u.unscoped
}

func (u *Unscoped) ApplyDeleteMany(options *DeleteManyOptions) {
	options.Unscoped = u.unscoped
}

func (u *Unscoped) ApplyFind(options *FindOptions) {
	options.Unscoped = u.unscoped
}

func (u *Unscoped) ApplyFindOne(options *FindOneOptions) {
	options.Unscoped = u.unscoped
}

func (u *Unscoped) ApplyReplaceOne(options *ReplaceOneOptions) {
	options.Unscoped = u.unscoped
}

func (u *Unscoped) ApplyUpdateOne(options *UpdateOneOptions) {
	options.Unscoped = u.unscoped
}

func (u *Unscoped) ApplyUpdateMany(options *UpdateManyOptions) {
	options.Unscoped = u.unscoped
}
