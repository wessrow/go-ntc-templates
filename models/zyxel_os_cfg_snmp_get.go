package models

type ZyxelOsCfgSnmpGet struct {
	Agent	string	`json:"AGENT"`
	Get_community	string	`json:"GET_COMMUNITY"`
	Set_community	string	`json:"SET_COMMUNITY"`
	Trap_community	string	`json:"TRAP_COMMUNITY"`
	System_name	string	`json:"SYSTEM_NAME"`
	System_location	string	`json:"SYSTEM_LOCATION"`
	System_contact	string	`json:"SYSTEM_CONTACT"`
	Domain_name	string	`json:"DOMAIN_NAME"`
	Trap_dest	string	`json:"TRAP_DEST"`
}

var ZyxelOsCfgSnmpGet_Template = `Value AGENT (\S*)
Value GET_COMMUNITY (\S*)
Value SET_COMMUNITY (\S*)
Value TRAP_COMMUNITY (\S*)
Value SYSTEM_NAME (\S*)
Value SYSTEM_LOCATION (.*)
Value SYSTEM_CONTACT (.*)
Value DOMAIN_NAME (\S*)
Value TRAP_DEST (.*)

Start
  ^SNMP\sAgent\s*:\s*${AGENT}
  ^Get\sCommunity\s*:\s*${GET_COMMUNITY}
  ^Set\sCommunity\s*:\s*${SET_COMMUNITY}
  ^Trap\sCommunity\s*:\s*${TRAP_COMMUNITY}
  ^System\sName\s*:\s*${SYSTEM_NAME}
  ^System\sLocation\s*:\s*${SYSTEM_LOCATION}
  ^System\sContact\s*:\s*${SYSTEM_CONTACT}
  ^Domain\sName\s*:\s*${DOMAIN_NAME}
  ^Trap\sDestination\s*:\s*${TRAP_DEST}
  ^Command Successful.
  ^\s*$$
  ^. -> Error

`