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

var CiscoNxosShowFexId_Template = `Value FEX (\d+)
Value DESCRIPTION (\S+)
Value STATE (\S+)
Value VERSION (\S+)
Value EXTENDER_SERIAL (\S+)
Value EXTENDER_MODEL (\S+)
Value PART_NO (\S+)
Value PINNING_MODE (\w+)
Value MAX_LINKS (\d+)
Value FABRIC_PORT (\S+)
Value FCOE_ADMIN (\w+)
Value FCOE_OPER (\w+)
Value FCOE_FEX_AA_CONFIGURED (\w+)
Value List FABRIC_INTERFACE (\S+)
Value List INTERFACE_MODE (\w+)
Value List INTERFACE_STATE (\w+)

Start
  ^FEX:\s+${FEX}\s+Description:\s+${DESCRIPTION}\s+state:\s+${STATE}
  ^.*FEX\s+version:\s+${VERSION}
  ^.*Extender\s+Serial:\s+${EXTENDER_SERIAL}
  ^.*Extender\s+Model:\s+${EXTENDER_MODEL},.*Part\s+No:\s+${PART_NO}
  ^.*Pinning\-mode:\s+${PINNING_MODE}.*Max\-links:\s+${MAX_LINKS}
  ^.*Fabric\s+port\s+for\s+control\s+traffic:\s+${FABRIC_PORT} 
  ^.*FCoE\s+Admin:\s+${FCOE_ADMIN}
  ^.*FCoE\s+Oper:\s+${FCOE_OPER}  
  ^.*FCoE\s+FEX\s+AA\s+Configured:\s+${FCOE_FEX_AA_CONFIGURED} 
  ^\s+${FABRIC_INTERFACE}\s+\-\s+Interface\s+${INTERFACE_MODE}\.\s+State:\s+${INTERFACE_STATE} 
  ^\S -> Record

`