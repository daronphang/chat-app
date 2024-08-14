package domain

type UserContact struct {
	UserID 		string 		`json:"userId" validate:"required"`
	Email		string 		`json:"email" validate:"required"`
}