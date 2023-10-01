package models

type CreateMovieDTO struct {
	Name        string            `json:"name" validate:"required,min=0,max=60"`
	ReleaseDate string            `json:"date" validate:"required,datetime=2006-01-02"`
	Genre       string            `json:"genre" validate:"required,oneof=Comedy Thriller Horror Drama Melody Melo-drama"`
	Director    CreateDirectorDTO `json:"director" validate:"required"`
}
type CreateDirectorDTO struct {
	Name   string `json:"name" validate:"required,min=0,max=60"`
	Gender string `json:"gender" validate:"required,oneof=Male Female"`
	Email  string `json:"email" validate:"required,email"`
}
