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
