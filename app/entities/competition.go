package entities

type Competition struct {
	UID  string
	Year string
}

type CompetitionFilter struct {
	UID  *string
	Year *string
}
