package models

type HuaweiSmartaxDisplayServicePortAll struct {
	Index	string	`json:"INDEX"`
	Vlan_id	string	`json:"VLAN_ID"`
	Vlan_attr	string	`json:"VLAN_ATTR"`
	Port_type	string	`json:"PORT_TYPE"`
	Fsp	string	`json:"FSP"`
	Ont_id	string	`json:"ONT_ID"`
	Gem_index	string	`json:"GEM_INDEX"`
	Flow_type	string	`json:"FLOW_TYPE"`
	Vlan_access	string	`json:"VLAN_ACCESS"`
	Rx	string	`json:"RX"`
	Tx	string	`json:"TX"`
	State	string	`json:"STATE"`
}

var HuaweiSmartaxDisplayServicePortAll_Template = `Value INDEX (\d+)
Value VLAN_ID (\d+)
Value VLAN_ATTR (\S+)
Value PORT_TYPE (\S+)
Value FSP (\d+\/\d+\s*\/\d+)
Value ONT_ID (\d+)
Value GEM_INDEX (\d+)
Value FLOW_TYPE (\S+)
Value VLAN_ACCESS (\d+)
Value RX (\S+)
Value TX (\S+)
Value STATE (\S+)


Start
  ^Switch-Oriented\s*Flow\s*List
  ^\s*-
  ^\s*INDEX\s+VLAN\s+VLAN\s+PORT\s+F\/\s*S\/\s*P\s*VPI\s*VCI\s*FLOW\s*FLOW\s*RX\s*TX\s*STATE\s*$$
  ^\s*ID\s*ATTR\s*TYPE\s*TYPE\s*PARA\s*$$ -> ServicePorts
  ^\s*$$
  ^. -> Error

ServicePorts
  ^\s*${INDEX}\s*${VLAN_ID}\s*${VLAN_ATTR}\s*${PORT_TYPE}\s*${FSP}\s*${ONT_ID}\s*${GEM_INDEX}\s*${FLOW_TYPE}\s*${VLAN_ACCESS}\s*(-|${RX})\s*(-|${TX})\s*${STATE}\s* -> Record
  ^\s*.
  ^\s*$$
  ^. -> Error

`