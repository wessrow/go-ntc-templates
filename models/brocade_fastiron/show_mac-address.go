package brocade_fastiron

type ShowMacAddress struct {
	Mac_address      string `json:"MAC_ADDRESS"`
	Destination_port string `json:"DESTINATION_PORT"`
	Age              string `json:"AGE"`
	Vlan_id          string `json:"VLAN_ID"`
	Type             string `json:"TYPE"`
}

var ShowMacAddress_Template string = `Value Required MAC_ADDRESS ([A-Fa-f0-9\.]{14})
Value Required DESTINATION_PORT (\S+)
Value Required AGE (\d+)
Value Required VLAN_ID (\d+)
Value TYPE (\S+)


Start
  ^Total.*
  ^MAC Address\s+Port\s+Age\s+VLAN\s+Type
  ^${MAC_ADDRESS}\s+${DESTINATION_PORT}\s+${AGE}\s+${VLAN_ID}\s+${TYPE} -> Record
  ^${MAC_ADDRESS}\s+${DESTINATION_PORT}\s+${AGE}\s+${VLAN_ID} -> Record
  ^. -> Error`
