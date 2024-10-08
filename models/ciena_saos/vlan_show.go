package ciena_saos 

type VlanShow struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Vlan_name	string	`json:"VLAN_NAME"`
}

var VlanShow_Template = `Value VLAN_ID (\d+)
Value VLAN_NAME ([0-9a-zA-Z_#]+)

Start
  ^\+-+\s*VLAN\s*GLOBAL\s*CONFIGURATION\s*-+\+
  ^\+-+\+-+.*
  ^\|\sParameter\s+|\sValue.*
  ^.*Inner TPID State.*
  ^\s+\|\s+\|\s+\|
  ^\|\s+Field\s+\|\s+Admin\s+\|\s+Oper\s+\|
  ^\|\s+Inner TPID.*
  ^\|\s+\|\s+\|
  ^\|VLAN\|.*
  ^\|\s+ID\s+\|\s+VLAN Name* -> VLAN_TABLE
  ^\+-+\s+CROSS\s+CONNECTION\s+TABLE -> CROSS_TABLE
  ^\s*$$
  ^. -> Error

VLAN_TABLE
  ^\|\s*${VLAN_ID}\s*\|\s*${VLAN_NAME}\s*\| -> Record
  ^\+-+\+.*
  ^\+-+\s+CROSS\s+CONNECTION\s+TABLE -> CROSS_TABLE
  ^\s*$$
  ^. -> Error
  
CROSS_TABLE
  ^\|\s+OVID
  ^\+-+\+.*
  ^\|\s+No Entry Found
  ^\s*$$
  ^. -> Error
`