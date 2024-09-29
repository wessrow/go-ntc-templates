package huawei_smartax 

type DisplayMacAddress struct {
	Service_port_index	string	`json:"SERVICE_PORT_INDEX"`
	Bundle_index	string	`json:"BUNDLE_INDEX"`
	Type	string	`json:"TYPE"`
	Mac	string	`json:"MAC"`
	Mac_type	string	`json:"MAC_TYPE"`
	Fsp	string	`json:"FSP"`
	Ont_id	string	`json:"ONT_ID"`
	Gem_index	string	`json:"GEM_INDEX"`
	Vlan_id	string	`json:"VLAN_ID"`
}

var DisplayMacAddress_Template = `Value SERVICE_PORT_INDEX (\d+)
Value BUNDLE_INDEX (\d+)
Value TYPE (eth|gpon|epon)
Value Required MAC ([0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4})
Value MAC_TYPE (static|dynamic)
Value FSP (\d+\s*\/\d+\s*\/\d+)
Value ONT_ID (\d+)
Value GEM_INDEX (\d+)
Value VLAN_ID (\d+)


Start
  ^\s*It\s*will\s*take\s*some\s*time,\s*please\s*wait...
  ^\s*-
  ^\s*SRV-P\s*BUNDLE\s*TYPE\s*MAC\s*MAC\s*TYPE\s*F\s*\/S\s*\/P\s*VPI\s*VCI\s*VLAN\s*ID\s*$$
  ^\s*INDEX\s*INDEX\s*$$ -> MACs
  ^\s*$$
  ^. -> Error

MACs
  ^\s*(-|${SERVICE_PORT_INDEX})\s*(-|${BUNDLE_INDEX})\s*${TYPE}\s*${MAC}\s*${MAC_TYPE}\s*${FSP}\s*(-|${ONT_ID})\s*(-|${GEM_INDEX})\s*${VLAN_ID} -> Record
  ^\s*-
  ^\s*Total:\s*\d+ -> EOF
  ^\s*$$
  ^. -> Error
`