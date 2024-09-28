package models

type AvayaErsShowMltAllMembers struct {
	Mltid	string	`json:"MLTID"`
	Mltname	string	`json:"MLTNAME"`
	Mltactivemembers	string	`json:"MLTACTIVEMEMBERS"`
	Mltinactivemembers	string	`json:"MLTINACTIVEMEMBERS"`
	Mltbpdu	string	`json:"MLTBPDU"`
	Mltmode	string	`json:"MLTMODE"`
	Mltstatus	string	`json:"MLTSTATUS"`
	Mlttype	string	`json:"MLTTYPE"`
	Mltlacpkey	string	`json:"MLTLACPKEY"`
}