package log

import "time"

type LogCore struct {
	ID        int
	Feedback  string
	Status_id int
	Mentee_id Mentee
	User_id   User
	Proof     []Proof
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Proof struct {
	LogID uint
	Link  string
}

type Status struct {
	ID   int
	Name string
}

type Mentee struct {
	ID   int
	Name string
}

type User struct {
	ID   int
	Name string
}
