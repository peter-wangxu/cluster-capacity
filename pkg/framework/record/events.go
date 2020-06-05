package record

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
)

type EventRecorder struct {
	Events chan Event
}

func (er *EventRecorder) Eventf(regarding runtime.Object, related runtime.Object, eventtype, reason, action, note string, args ...interface{}) {
	if er.Events != nil {
		er.Events <- Event{eventtype, reason, fmt.Sprintf(action, args...)}
	}
}

// NewFakeRecorder creates new fake event recorder with event channel with
// buffer of given size.
func NewEventRecorder(bufferSize int) *EventRecorder {
	return &EventRecorder{
		Events: make(chan Event, bufferSize),
	}
}
