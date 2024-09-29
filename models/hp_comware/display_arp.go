package hp_comware 

type DisplayArp struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Vlan_id	string	`json:"VLAN_ID"`
	Interface	string	`json:"INTERFACE"`
	Aging	string	`json:"AGING"`
	Type	string	`json:"TYPE"`
}

var DisplayArp_Template = `Value IP_ADDRESS (\d+.\d+.\d+.\d+)
Value MAC_ADDRESS (\w+-\w+-\w+)
Value VLAN_ID (\S+|\d+)
Value INTERFACE (\S+)
Value AGING (\d+)
Value TYPE (\S+)

Start
  ^${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${VLAN_ID}\s+${INTERFACE}\s+${AGING}\s+${TYPE} -> Record
`