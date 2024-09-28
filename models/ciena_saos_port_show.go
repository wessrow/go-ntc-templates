package models

type CienaSaosPortShow struct {
	Name	string	`json:"NAME"`
	Type	string	`json:"TYPE"`
	Link	string	`json:"LINK"`
	Duration	string	`json:"DURATION"`
	Xcvr	string	`json:"XCVR"`
	Stp	string	`json:"STP"`
	Mode	string	`json:"MODE"`
	Autoneg	string	`json:"AUTONEG"`
	Admin_link	string	`json:"ADMIN_LINK"`
	Admin_mode	string	`json:"ADMIN_MODE"`
	Admin_autoneg	string	`json:"ADMIN_AUTONEG"`
}