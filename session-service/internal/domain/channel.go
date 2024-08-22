package domain

type Channel struct {
	ChannelID 	string		`json:"channelId" validate:"required"` 
	ChannelName string		`json:"channelName" validate:"required"`
	CreatedAt 	string 		`json:"createdAt" validate:"required"`
	UserIDs 	[]string	`json:"userIds" validate:"required"`
}
