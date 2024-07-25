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

type UserStatus struct {
	UserID		string
	TargetID	string
	Status		string
}