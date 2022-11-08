package events

import (
	"errors"
	"sync"
)

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
	if handlers, ok := eventDispatcher.handlers[eventName]; ok {
		for _, eventHandler := range handlers {
			if eventHandler == handler {

				return true
			}
		}
	}

	return false
}

func (eventDispatcher *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := eventDispatcher.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}

		for _, eventHandler := range handlers {
			wg.Add(1)
			go eventHandler.Handle(event, wg)
		}
		wg.Wait()

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
