package domain

type UserSession struct {
	UserID 			string 
	Servers			[]string
	LastHeartbeat 	string
}

type HeartbeatRequest struct {
	UserID string
	Server string
}
