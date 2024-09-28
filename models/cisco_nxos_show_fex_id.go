package models

type CiscoNxosShowFexId struct {
	Fex	string	`json:"FEX"`
	Description	string	`json:"DESCRIPTION"`
	State	string	`json:"STATE"`
	Version	string	`json:"VERSION"`
	Extender_serial	string	`json:"EXTENDER_SERIAL"`
	Extender_model	string	`json:"EXTENDER_MODEL"`
	Part_no	string	`json:"PART_NO"`
	Pinning_mode	string	`json:"PINNING_MODE"`
	Max_links	string	`json:"MAX_LINKS"`
	Fabric_port	string	`json:"FABRIC_PORT"`
	Fcoe_admin	string	`json:"FCOE_ADMIN"`
	Fcoe_oper	string	`json:"FCOE_OPER"`
	Fcoe_fex_aa_configured	string	`json:"FCOE_FEX_AA_CONFIGURED"`
	Fabric_interface	[]string	`json:"FABRIC_INTERFACE"`
	Interface_mode	[]string	`json:"INTERFACE_MODE"`
	Interface_state	[]string	`json:"INTERFACE_STATE"`
}