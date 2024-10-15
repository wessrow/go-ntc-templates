package juniper_junos

type ShowEthernetSwitchingTable struct {
	Vlan_id           string `json:"VLAN_ID"`
	Bridge_domain_id  string `json:"BRIDGE_DOMAIN_ID"`
	Mac_address       string `json:"MAC_ADDRESS"`
	Mac_flag          string `json:"MAC_FLAG"`
	Age               string `json:"AGE"`
	Logical_interface string `json:"LOGICAL_INTERFACE"`
}

var ShowEthernetSwitchingTable_Template string = `Value VLAN_ID (\S+)
Value BRIDGE_DOMAIN_ID (\S+)
Value MAC_ADDRESS (\S+)
Value MAC_FLAG (\S+)
Value AGE (\S+)
Value LOGICAL_INTERFACE (\S+)

Start
  ^MAC\sflags
  ^\s*SE\s-\sstatistics\senabled
  ^Ethernet\sswitching\stable\s:\s\d+\sentries,\s\d+\slearned$$
  ^Routing\sinstance\s:\sdefault-switch$$
  ^\s*Vlan\s+MAC\s+MAC\s+Age\s+Logical\s*\S*\s*\S*\s*$$
  ^\s*name\s+address\s+flags\s+interface\s*\S*\s*\S*\s*$$
  ^\s*(?:Vlan${VLAN_ID}|BD-${BRIDGE_DOMAIN_ID})\s+${MAC_ADDRESS}\s+${MAC_FLAG}\s+${AGE}\s+${LOGICAL_INTERFACE}\s*\S*\s*\S* -> Record
  ^\s*$$
  ^{master:\d+}
  ^. -> Error
`
