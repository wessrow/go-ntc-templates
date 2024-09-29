package ciena_saos 

type PortXcvrShow struct {
	Port	string	`json:"PORT"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Operational_state	string	`json:"OPERATIONAL_STATE"`
	Vendor	string	`json:"VENDOR"`
	Ciena_rev	string	`json:"CIENA_REV"`
	Connector_type	string	`json:"CONNECTOR_TYPE"`
	Diagnostic	string	`json:"DIAGNOSTIC"`
}

var PortXcvrShow_Template = `Value PORT (\S+)
Value ADMIN_STATE (\S+)
Value OPERATIONAL_STATE (\S+)
Value VENDOR (\S.*\S)
Value CIENA_REV (\S+)
Value CONNECTOR_TYPE (\S+)
Value DIAGNOSTIC (\S+)

Start
  ^\+-+
  ^\|\s+\|\s*Admin\s*\|\s*Oper\s*\|\s+\|\s*Ciena\s*\|\s*Ether\s*Medium\s*&\s*\|\s*Diag\s*\|\s*$$
  ^\|\s*Port\s*\|(?:\s*State\s*\|){2}\s*Vendor\s*Name\s*&\s*Part\s*Number\s*\|\s*Rev\s*\|\s*Connector\s*Type\s*\|\s*Data\s*\|\s*$$
  ^\|\s*${PORT}\s*\|${ADMIN_STATE}\s*\|(${OPERATIONAL_STATE}|)\s*\|(${VENDOR}|)\s*\|(${CIENA_REV}|)\s*\|(${CONNECTOR_TYPE}|)\s*\|(${DIAGNOSTIC}|)\s*\| -> Record
  ^\s*$$
  ^. -> Error
`