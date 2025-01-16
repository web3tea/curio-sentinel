package processor

import (
	"github.com/strahe/curio-sentinel/capturer"
)

type EventProcessor interface {
	Process(event *capturer.Event) (*capturer.Event, error)
}

type ProcessorComposite interface {
	AddFilter(processor EventProcessor)
	AddTransformer(processor EventProcessor)
}

type Processor interface {
	EventProcessor
	ProcessorComposite
}
