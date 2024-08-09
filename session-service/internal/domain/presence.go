package domain

type User struct {
	UserID 			string 
	Servers			[]string
	LastHeartbeat 	string
}

type HeartbeatRequest struct {
	UserID string
	Server string
}

type UserPresence struct {
	UserID		string		`json:"userId" validate:"required"`
	TargetIDs	[]string 	`json:"targetIds" validate:"required"`
	Status		string 		`json:"status" validate:"required,oneof=online offline"` 
}