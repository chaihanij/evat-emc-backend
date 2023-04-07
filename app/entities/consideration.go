package entities

type AssignmentScore struct {
	// ID string
	// UpdatedAt     time.Time
	// No            int
	// TotalScore    float64
	Total          float64
	Considerations []Consideration
}

type Consideration struct {
	ID       string
	TeamName string
	Score    float64
	Title    string
}
type ConsiderationFilter struct {
	AssignmentUUID *string
	ID             *string
}
