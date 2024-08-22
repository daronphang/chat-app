package domain

type ServerMetadata struct {
	Name 		string 	`json:"name"`
	URL 		string	`json:"url"`
	CPU 		float32	`json:"cpu"`
	Memory 		float32	`json:"memory"`
}