package domain

type PresenceStatus struct {
	TargetID	string 	`json:"targetId" validate:"required"` 
	ClientID 	string 	`json:"clientId" validate:"required"` 
	Status		string	`json:"status" validate:"required,oneof=online offline"` 
}