package govent

type Payload = map[string]interface{}
type EventHandler = func(data Payload)

type IEventEmitter interface {
	On(eventName string, handler EventHandler)
	Emit(eventName string, data Payload) bool
}

type EventEmitter struct {
	events map[string]EventHandler
}

func (e *EventEmitter) On(eventName string, handler EventHandler) {
	e.events[eventName] = handler
}

func (e *EventEmitter) Emit(eventName string, data Payload) bool {
	if handler, ok := e.events[eventName]; !ok {
		return false
	} else {
		handler(data)
		return true
	}
}

func New() *EventEmitter {
	return &EventEmitter{
		events: make(map[string]EventHandler),
	}
}
