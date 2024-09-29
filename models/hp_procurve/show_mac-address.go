package hp_procurve 

type ShowMacAddress struct {
	Mac_address	string	`json:"MAC_ADDRESS"`
	Port	string	`json:"PORT"`
	Vlan_id	string	`json:"VLAN_ID"`
}

var ShowMacAddress_Template = `Value MAC_ADDRESS ([0-9a-fA-F]{6}-[0-9a-fA-F]{6})
Value PORT (\S+)
Value VLAN_ID (\d+)

Start
  ^\s*--- -> Start_record

Start_record
  ^\s+${MAC_ADDRESS}\s+${PORT}\s+${VLAN_ID} -> Record
  ^\s+${MAC_ADDRESS}\s+${PORT} -> Record
  ^\s*$$
  ^. -> Error
`