package models

type Medication struct {
	AppModel
	Name     string
	Weight   int
	Code     string
	ImageURL string
}
