package models

type Project struct {
	ID               int
	Name             string
	Description      string
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
