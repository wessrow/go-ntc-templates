package ruckus_fastiron

type ShowMacAddress struct {
	Mac_address string `json:"MAC_ADDRESS"`
	Port        string `json:"PORT"`
	Type        string `json:"TYPE"`
	Vlan_id     string `json:"VLAN_ID"`
}

var ShowMacAddress_Template string = `Value MAC_ADDRESS ([A-Fa-f0-9\.]{14})
Value PORT (\S+)
Value TYPE (\S+)
Value VLAN_ID ([0-9]*)


Start
  ^Total\s+active\s+entries\s+from\s+all\s+ports\s+=\s+[0-9]*
  ^MAC-Address\s+Port\s+Type\s+VLAN
  ^${MAC_ADDRESS}\s+${PORT}\s+${TYPE}\s+${VLAN_ID} -> Record
  ^\s*$$
  ^. -> Error
`
