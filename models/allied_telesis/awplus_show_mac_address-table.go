package allied_telesis

type AwplusShowMacAddressTable struct {
	Vlan_id             string `json:"VLAN_ID"`
	Destination_port    string `json:"DESTINATION_PORT"`
	Destination_address string `json:"DESTINATION_ADDRESS"`
	Fwd                 string `json:"FWD"`
}

var AwplusShowMacAddressTable_Template string = `Value Required VLAN_ID (\d+)
Value DESTINATION_PORT (\S+)
Value DESTINATION_ADDRESS (\S+)
Value FWD (\S+)

Start
  ^VLAN\s+port
  ^${VLAN_ID}\s+${DESTINATION_PORT}\s+${DESTINATION_ADDRESS}\s+${FWD} -> Record
  ^. -> Error
`
