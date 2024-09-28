package models

type CienaSaosPortXcvrShow struct {
	Port	string	`json:"PORT"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Operational_state	string	`json:"OPERATIONAL_STATE"`
	Vendor	string	`json:"VENDOR"`
	Ciena_rev	string	`json:"CIENA_REV"`
	Connector_type	string	`json:"CONNECTOR_TYPE"`
	Diagnostic	string	`json:"DIAGNOSTIC"`
}