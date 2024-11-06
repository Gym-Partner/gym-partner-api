package model

type Serie struct {
	Id          string `json:"id" gorm:"primaryKey, not null"`
	Weight      int    `json:"weight"`
	Repetitions int    `json:"repetitions"`
	IsWarmUp    bool   `json:"is_warm_up"`
}

type Series []Serie
