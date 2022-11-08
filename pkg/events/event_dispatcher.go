package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (eventDispatcher *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := eventDispatcher.handlers[eventName]; ok {
		for _, eventHandler := range eventDispatcher.handlers[eventName] {
			if eventHandler == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	eventDispatcher.handlers[eventName] = append(eventDispatcher.handlers[eventName], handler)
	return nil
}

func (eventDispatcher *EventDispatcher) Clear() error {
	eventDispatcher.handlers = make(map[string][]EventHandlerInterface)
	return nil
}

func (eventDispatcher *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := eventDispatcher.handlers[eventName]; ok {
		for _, eventHandler := range eventDispatcher.handlers[eventName] {
			if eventHandler == handler {
				return true
			}
		}
	}

	return false
}

func (eventDispatcher *EventDispatcher) Dispatch(event EventInterface) error {
	if _, ok := eventDispatcher.handlers[event.GetName()]; ok {
		for _, eventHandler := range eventDispatcher.handlers[event.GetName()] {
			eventHandler.Handle(event)
		}
	}

	return nil
}

func (eventDispatcher *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := eventDispatcher.handlers[eventName]; ok {
		for i, eventHandler := range eventDispatcher.handlers[eventName] {
			if eventHandler == handler {
				eventDispatcher.handlers[eventName] = append(eventDispatcher.handlers[eventName][:i], eventDispatcher.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}

	return nil
}
