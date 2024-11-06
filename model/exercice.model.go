package model

type Exercice struct {
	Id         string `json:"id" gorm:"primaryKey, not null"`
	Name       string `json:"name"`
	Equipement bool   `json:"equipement"`
}

type Exercices []Exercice
