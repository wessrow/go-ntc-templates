package models

type CiscoIosShowVlan struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Vlan_name	string	`json:"VLAN_NAME"`
	Status	string	`json:"STATUS"`
	Interfaces	[]string	`json:"INTERFACES"`
}