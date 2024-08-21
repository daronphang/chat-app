package domain

type ServerMetadata struct {
	Name 		string 	`json:"name"`
	URL 		string	`json:"url"`
	CPU 		float64	`json:"cpu"`
	Memory 		float64	`json:"memory"`
}