package domain

type NewUser struct {
	UserID 		string 	
	Email		string 	`json:"email" validate:"required"`
	DisplayName	string	`json:"displayName" validate:"required"`
}

type UserMetadata struct {
	UserID 		string 		`json:"userId" validate:"required"`
	Email		string 		`json:"email" validate:"required"`
	DisplayName	string		`json:"displayName" validate:"required"`
	CreatedAt	string		`json:"createdAt"`
	Friends		[]Friend	`json:"friends"`
}

type UserContact struct {
	UserID 		string 		`json:"userId" validate:"required"`
	Email		string 		`json:"email" validate:"required"`
}

type UserCredentials struct {
	Email		string 	`json:"email" validate:"required"`
}

type NewFriend struct {
	UserID		string 	`json:"userId"`
	FriendEmail string	`json:"friendEmail" validate:"required"`
	FriendID 	string	`json:"friendId"`
	DisplayName string	`json:"displayName" validate:"required"`
}

type Friend struct {
	UserID		string 	`json:"userId" validate:"required"`
	Email 		string	`json:"email" validate:"required"`
	DisplayName string	`json:"displayName" validate:"required"`
}

