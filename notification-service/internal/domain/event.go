package domain

type Event string

var (
	EventChannel 		Event = "event/channel"
	EventMessage 		Event = "event/Message"
	EventUserPresence 	Event = "event/presence"
)

type BaseEvent struct {
	Event 		Event 		`json:"event" validate:"required"`
	Timestamp	string		`json:"timestamp" validate:"required"`
	Data 		interface{} `json:"data" validate:"required"`
}

