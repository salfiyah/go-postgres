package models

// User Schema of the user table

type User struct {
	ID			int64	`json:"id"`
	Name		string  `Json:"name"`
	Location	string  `json:"location"`
	Age			int64	`json:"age"`
}


