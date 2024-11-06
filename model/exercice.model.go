package model

type Exercice struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Equipement bool   `json:"equipement"`
}

type Exercices []Exercice
