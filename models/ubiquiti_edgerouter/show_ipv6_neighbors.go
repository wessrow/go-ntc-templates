package ubiquiti_edgerouter 

type ShowIpv6Neighbors struct {
	Ipv6_address	string	`json:"IPV6_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	State	string	`json:"STATE"`
}

var ShowIpv6Neighbors_Template = `Value Required IPV6_ADDRESS (\S+)
Value INTERFACE (\S+)
Value MAC_ADDRESS ((?:[0-9a-fA-F]{2}\:){5}[0-9a-fA-F]{2})
Value STATE (\S+)

Start
  ^${IPV6_ADDRESS}\s+(dev\s${INTERFACE})?\s*(lladdr\s${MAC_ADDRESS})?\s*(router\s)?${STATE}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`