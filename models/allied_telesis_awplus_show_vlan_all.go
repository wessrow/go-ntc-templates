package models

type AlliedTelesisAwplusShowVlanAll struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Name	string	`json:"NAME"`
	Status	string	`json:"STATUS"`
	Interfaces	[]string	`json:"INTERFACES"`
}