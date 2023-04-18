package entities

type AssignmentScore struct {
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

type FieldRaceTeamScore struct {
	Total          float64
	Considerations []Consideration
}

type AllScore struct {
	ID                string
	Name              string
	Allconsiderations []AllConsideration
	Total             float64
}
type AllConsideration struct {
	Type  string
	Score float64
	Title string
}

type Allconsiderations struct {
	Title string
	Score float64
	Type  string
}

type AllScoreFilter struct {
	Name string
}
