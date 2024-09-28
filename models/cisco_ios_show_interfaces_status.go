package models

type CiscoIosShowInterfacesStatus struct {
	Port	string	`json:"PORT"`
	Name	string	`json:"NAME"`
	Status	string	`json:"STATUS"`
	Vlan_id	string	`json:"VLAN_ID"`
	Duplex	string	`json:"DUPLEX"`
	Speed	string	`json:"SPEED"`
	Type	string	`json:"TYPE"`
	Fc_mode	string	`json:"FC_MODE"`
}