package domain

type NewUser struct {
	UserID 		string 	
	Email		string 	`json:"email" validate:"required"`
	DisplayName	string	`json:"displayName" validate:"required"`
}

type UserMetadata struct {
	UserID 		string 	`json:"userId" validate:"required"`
	Email		string 	`json:"email" validate:"required"`
	DisplayName	string	`json:"displayName" validate:"required"`
	CreatedAt	string	`json:"createdAt"`
}

type Login struct {
	Email		string 	`json:"email" validate:"required"`
}

type NewContact struct {
	UserID		string 	`json:"userId" validate:"required"`
	FriendEmail string	`json:"friendEmail" validate:"required"`
	FriendID 	string	
	DisplayName string	`json:"displayName" validate:"required"`
}

type Contact struct {
	UserID		string 	`json:"userId" validate:"required"`
	Email 		string	`json:"email" validate:"required"`
	DisplayName string	`json:"displayName" validate:"required"`
}