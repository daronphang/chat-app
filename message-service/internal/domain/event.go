package domain

type Event string

var (
	EventNewChannel 	Event = "event/channel/new"
	EventMessage 		Event = "event/Message"
	EventUserPresence 	Event = "event/presence"
)

type BaseEvent struct {
	Event 	Event 		`json:"event" validate:"required"`
	Data 	interface{} `json:"data" validate:"required"`
}