package model

type Serie struct {
	Id          int  `json:"id"`
	Weight      int  `json:"weight"`
	Repetitions int  `json:"repetitions"`
	IsWarmUp    bool `json:"is_warm_up"`
}

type Series []Serie
