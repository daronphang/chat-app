package domain

type Event string

var (
	EventChannel 		Event = "event/channel"
	EventMessage 		Event = "event/message"
	EventUserPresence 	Event = "event/presence"
)

type BaseEvent struct {
	Event 			Event 		`json:"event" validate:"required"`
	EventTimestamp	string		`json:"eventTimestamp" validate:"required"`
	Data 			interface{} `json:"data" validate:"required"`
}

