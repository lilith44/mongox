package timer

type Option interface {
	apply(*Timer)
}
