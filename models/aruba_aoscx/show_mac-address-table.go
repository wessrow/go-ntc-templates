package aruba_aoscx

type ShowMacAddressTable struct {
	Type        string `json:"TYPE"`
	Port        string `json:"PORT"`
	Mac_address string `json:"MAC_ADDRESS"`
	Vlan_id     string `json:"VLAN_ID"`
}

var ShowMacAddressTable_Template string = `Value MAC_ADDRESS (\S+)
Value VLAN_ID (\d+)
Value TYPE (\S+)
Value PORT (\S+)

Start
  ^MAC\s+age-time.*$$
  ^Number\s+of\s+MAC.*$$
  ^MAC\s+Address\s+VLAN\s+Type\s+Port
  ^-+$$
  ^${MAC_ADDRESS}\s+${VLAN_ID}\s+${TYPE}\s+${PORT} -> Record
  ^\s*$$
  ^. -> Error
`
