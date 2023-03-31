package entities

import "time"

// type EmailContact struct {
// 	ID          string
// 	Title string
// 	Email       string
// 	FirstName   string
// 	LastName    string
// 	Description string
// 	Create_at   time.Time
// 	Status      bool
// }

type CreateContactEmail struct {
	ID          string
	Title       string
	Email       string
	FirstName   string
	LastName    string
	Description string
	Create_at   time.Time
	Status      bool
}
