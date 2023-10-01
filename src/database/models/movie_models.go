package models

type Movie struct {
	Id          string
	Name        string
	ReleaseDate string
	Genre       string
	Director    Director
}
type Director struct {
	Name   string
	Gender string
	Email  string
}
