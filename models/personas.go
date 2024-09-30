package models

type Persona struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Story string `json:"story"`
}

var Personas []Persona
