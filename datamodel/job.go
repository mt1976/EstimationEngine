package datamodel

type JobDefinition struct {
	ID          string
	Name        string
	Period      string
	Description string
	Type        string
}

type JobList struct {
	Job []JobDefinition
}
