package models

type UpdateMovieDTO struct {
	Name        string            `json:"name" validate:"omitempty,max=60"`
	ReleaseDate string            `json:"date" validate:"omitempty,datetime=2006-01-02"`
	Genre       string            `json:"genre" validate:"omitempty,oneof=comedy thriller horror drama melody melo-drama"`
	Director    UpdateDirectorDTO `json:"director" validate:"omitempty"`
}
type UpdateDirectorDTO struct {
	Name   string `json:"name" validate:"omitempty,max=60"`
	Gender string `json:"gender" validate:"omitempty,oneof=male female"`
	Email  string `json:"email" validate:"omitempty,email"`
}
