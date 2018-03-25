package main

type EventStream struct {
	innerStream chan EventObject
	eventMap map[EventType][]Handler
}

func NewEventStream() *EventStream {
	return &EventStream{
		innerStream: make(chan EventObject),
		eventMap: make(map[EventType][]Handler),
	}
}

func (es *EventStream) Subscribe(handler Handler, eventType EventType) {
	if len(es.eventMap[eventType]) == 0 {
		es.eventMap[eventType] = make([]Handler, 0)
	}
	es.eventMap[eventType] = append(es.eventMap[eventType], handler)

	return
}

func (es *EventStream) Run() {
	for {
		e := <-es.innerStream
		es.notify(e.EventType, e.Event)
	}
}

func (es *EventStream) notify(eventType EventType, event Event) {
	for _, h := range es.eventMap[eventType] {
		h(event)
	}
}

func (es *EventStream) AddEvent(eventObject EventObject) {
	es.innerStream <- eventObject
}