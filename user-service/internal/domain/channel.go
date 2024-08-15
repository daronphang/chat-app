package domain

type Channel struct {
	ChannelID 		string		`json:"channelId"` 
	ChannelName		string		`json:"channelName"`
	CreatedAt 		string 		`json:"createdAt"`
	UserIDs 		[]string	`json:"userIds"`
	LastMessageID	uint64		`json:"lastMessageId"`
}

type NewChannel struct {
	UserIDs 		[]string	`json:"userIds" validate:"required"`
	ChannelName 	string		`json:"channelName" validate:"required"`
}

type GroupMembers struct {
	UserID			string 		`json:"userId" validate:"required"`
	UserIDs			[]string 	`json:"userIds" validate:"required"`
	ChannelID 		string		`json:"channelId" validate:"required"`
	LastMessageID 	uint64		`json:"lastMessageId"`
}

type AdminGroupMember struct {
	UserID		string 	`json:"userId" validate:"required"`
	ChannelID 	string	`json:"channelId" validate:"required"`
	ChannelName string	`json:"channelName"`
}

type LastReadMessage struct {
	ChannelID 		string		`json:"channelId" validate:"required"` 
	UserID 			string		`json:"userId" validate:"required"`
	LastMessageID	uint64		`json:"lastMessageId" validate:"required"`
}