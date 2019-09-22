package events

const (
	NormalEvent = iota
	OnceEvent
)

type Params map[string]interface{}
type eventCb func(params Params)
type event struct {
	_type int
	_cb   eventCb
}

func newEvent(eventType int, cb eventCb) event {
	return event{
		_type: eventType,
		_cb:   cb,
	}
}
