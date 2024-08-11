package domain

type Channel struct {
	ChannelID 	string	`json:"channelId"` 
	ChannelName string	`json:"channelName"`
	CreatedAt 	string 	`json:"createdAt"`
}

type NewChannel struct {
	ChannelID 		string		`json:"channelId"`
	UserIDs 		[]string	`json:"userIds" validate:"required"`
	ChannelName 	string		`json:"channelName"`
	CreatedAt 		string 		`json:"createdAt"`
}

type NewGroupMember struct {
	UserID		string 	`json:"userId" validate:"required"`
	ChannelID 	string	`json:"channelId" validate:"required"`
}