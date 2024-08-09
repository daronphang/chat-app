package domain

type NewChannelEvent struct {
	ClientID 		string 		`json:"targetClientId" validate:"required"`
	ChannelID 		string		`json:"channelId" validate:"required"` 
	ChannelName 	string		`json:"channelName" validate:"required"`
	CreatedAt 		string 		`json:"createdAt" validate:"required"`
}

type ChannelUserEvent struct {
	ClientID 		string 	`json:"targetClientId" validate:"required"`
	ChannelID 		string	`json:"channelId" validate:"required"`
	UserID 			string	`json:"userId" validate:"required"`
	Action 			string	`json:"action" validate:"required"`
	Email 			string	`json:"email" validate:"required"`
	DisplayName 	string	`json:"displayName" validate:"required"`
}
