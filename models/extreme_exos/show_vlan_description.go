package extreme_exos 

type ShowVlanDescription struct {
	Vlan_name	string	`json:"VLAN_NAME"`
	Vlan_id	string	`json:"VLAN_ID"`
	Description	string	`json:"DESCRIPTION"`
}

var ShowVlanDescription_Template = `Value VLAN_NAME (\S+)
Value VLAN_ID (\d+)
Value DESCRIPTION (.*)

Start
  ^(\s*=+)+\s*$$
  ^\s*Name\s+VID\s+Description\s*$$ -> VlanTable
  ^\s*
  ^. -> Error

VlanTable
  ^(\s*=+)+\s*$$
  ^\s*${VLAN_NAME}\s+${VLAN_ID}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*>\s+Indicates\s+description\s+string\s+truncated\s+past\s+\d+\s+characters\s*$$
  ^\s*Total\s+number\s+of\s+VLAN\(s\)\s*:\s*\d+\s*$$
  ^\s*
  ^. -> Error
`