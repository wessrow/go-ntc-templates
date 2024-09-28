package models

type AlcatelSrosShowLag struct {
	Lag_id	string	`json:"LAG_ID"`
	Adm	string	`json:"ADM"`
	Opr	string	`json:"OPR"`
	Weighted	string	`json:"WEIGHTED"`
	Threshold	string	`json:"THRESHOLD"`
	Up_count	string	`json:"UP_COUNT"`
	Mc_act_stdby	string	`json:"MC_ACT_STDBY"`
}