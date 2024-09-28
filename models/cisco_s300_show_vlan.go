package models

type CiscoS300ShowVlan struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Vlan_name	string	`json:"VLAN_NAME"`
	Interfaces	string	`json:"INTERFACES"`
	Type	string	`json:"TYPE"`
	Authorization	string	`json:"AUTHORIZATION"`
	Created_by	string	`json:"CREATED_BY"`
}