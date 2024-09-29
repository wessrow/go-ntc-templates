package paloalto_panos 

type ShowMacAll struct {
	Vlan_id	string	`json:"VLAN_ID"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
	Status	string	`json:"STATUS"`
	Ttl	string	`json:"TTL"`
}

var ShowMacAll_Template = `Value VLAN_ID (\S+)
Value MAC_ADDRESS ([0-9a-fA-F:]+)
Value INTERFACE (\S+)
Value STATUS (\S+)
Value TTL (\d+)

Start
  ^\s*--- -> Start_record

Start_record
  ^${VLAN_ID}\s+${MAC_ADDRESS}\s+${INTERFACE}\s+${STATUS}\s+${TTL} -> Record
  ^\s*$$
  ^. -> Error
`