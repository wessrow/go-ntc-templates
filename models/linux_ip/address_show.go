package linux_ip 

type AddressShow struct {
	Id	string	`json:"ID"`
	Interface	string	`json:"INTERFACE"`
	Flags	string	`json:"FLAGS"`
	Mtu	string	`json:"MTU"`
	Qdisc	string	`json:"QDISC"`
	State	string	`json:"STATE"`
	Group	string	`json:"GROUP"`
	Qlen	string	`json:"QLEN"`
	Master	string	`json:"MASTER"`
	Type	string	`json:"TYPE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Broadcast	string	`json:"BROADCAST"`
	Ip_addresses	[]string	`json:"IP_ADDRESSES"`
	Ip_masks	[]string	`json:"IP_MASKS"`
	Ipv6_addresses	[]string	`json:"IPV6_ADDRESSES"`
	Ipv6_masks	[]string	`json:"IPV6_MASKS"`
}

var AddressShow_Template = `Value Required ID (\d+)
Value Required INTERFACE ([^:]+)
Value Required FLAGS (\S+)
Value Required MTU (\d+)
Value Required QDISC (\S+)
Value Required STATE (\S+)
Value GROUP (\S+)
Value QLEN (\d+)
Value MASTER (\S+)
Value Required TYPE (\S+)
Value MAC_ADDRESS ((?:[a-fA-F0-9:]{17}))
Value BROADCAST ((?:[a-fA-F0-9:]{17}))
Value List IP_ADDRESSES ([0-9\.]+)
Value List IP_MASKS (\d{1,2})
Value List IPV6_ADDRESSES ([0-9A-Fa-f:]+)
Value List IPV6_MASKS (\d{1,3})

Start
  ^\d+: -> Continue.Record
  ^${ID}:\s+${INTERFACE}:\s+<${FLAGS}>\s+mtu\s+${MTU}\s+qdisc\s+${QDISC}(?:\s+master\s+${MASTER})?\s+state\s+${STATE}(?:\s+group\s+${GROUP})?(?:\s+qlen\s+${QLEN})?\s*$$
  ^\s+link/${TYPE}(?:\s+${MAC_ADDRESS}\s+brd\s+${BROADCAST})?.*$$
  ^\s+altname.*$$
  ^\s+inet\s+${IP_ADDRESSES}/${IP_MASKS}\s+(:?brd|scope).*$$
  ^\s+inet6\s+${IPV6_ADDRESSES}/${IPV6_MASKS}\s+(?:brd|scope).*$$
  ^\s+valid_lft.*$$
  ^\s*$$
  ^. -> Error
`