package app

import "github.com/asynkron/protoactor-go/cluster"

type EventEnvelope struct {
	Event interface{}
}

func (e *EventEnvelope) GetPriority() int8 {
	return 7
}

type EventPublisher struct{}

func (e *EventPublisher) PublishEvent(event interface{}, ctx cluster.GrainContext) {
	ctx.Send(ctx.Self(), &EventEnvelope{Event: event})
}
