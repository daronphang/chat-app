package domain


type UserPresenceEvent struct {
	TargetID	string 	`json:"targetId" validate:"required"` // UserID to notify.
	ClientID 	string 	`json:"clientId" validate:"required"`
	Status		string	`json:"status" validate:"required,oneof=online offline"` 
}
type UserPresence struct {
	UserID		string		`json:"userId" validate:"required"`
	Status		string 		`json:"status" validate:"required,oneof=online offline"` 
}