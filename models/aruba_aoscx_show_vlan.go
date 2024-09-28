package models

type ArubaAoscxShowVlan struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Vlan_name	string	`json:"VLAN_NAME"`
	Status	string	`json:"STATUS"`
	Reason	string	`json:"REASON"`
	Type	string	`json:"TYPE"`
	Interfaces	[]string	`json:"INTERFACES"`
}