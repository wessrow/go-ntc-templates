package cisco_s300

type ShowMacAddressTable struct {
	Destination_address string `json:"DESTINATION_ADDRESS"`
	Type                string `json:"TYPE"`
	Vlan_id             string `json:"VLAN_ID"`
	Destination_port    string `json:"DESTINATION_PORT"`
}

var ShowMacAddressTable_Template string = `Value DESTINATION_ADDRESS ((\w\w:){5}\w\w)
Value TYPE (dynamic|self)
Value VLAN_ID (\w+)
Value DESTINATION_PORT (\S+)

Start
  ^Flags
  ^Aging
  ^\s+Vlan\s+Mac\s+Address\s+Port\s+Type -> Begin
  ^\s*$$
  ^. -> Error

Begin
  ^-+
  ^\s+${VLAN_ID}\s+${DESTINATION_ADDRESS}\s+${DESTINATION_PORT}\s+${TYPE} -> Record
  ^\s*$$
  ^. -> Error
`
