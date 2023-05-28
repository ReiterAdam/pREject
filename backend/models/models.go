package models

type Project struct {
	ID               int
	Name             string
	Description      string
	Author           string
	CreatedOn        string
	ModifiedOn       string
	CustomProperties []Property
}

type Position struct {
	ID               int
	Title            string
	CustomProperties []Property
}

type Property struct {
	Key   string
	Value string
}
