package events

type EventEmitter struct {
	events map[string][]event
}

func (e *EventEmitter) On(name string, cb func(params Params)) {
	e.addListener(name, cb)
}

func (e *EventEmitter) Emit(name string) {

}

func (e *EventEmitter) addListener(name string, cb eventCb) {
	e.events[name] = append(e.events[name], newEvent(NormalEvent, cb))
}
