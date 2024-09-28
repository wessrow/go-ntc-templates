package models

type CienaSaosChassisShowTemperature struct {
	Current	string	`json:"CURRENT"`
	Low	string	`json:"LOW"`
	High	string	`json:"HIGH"`
}