package linux

type IpAddressShow struct {
	Mac_address    string   `json:"MAC_ADDRESS"`
	Ip_masks       []string `json:"IP_MASKS"`
	Ipv6_masks     []string `json:"IPV6_MASKS"`
	Master         string   `json:"MASTER"`
	Flags          string   `json:"FLAGS"`
	Group          string   `json:"GROUP"`
	Type           string   `json:"TYPE"`
	Ip_addresses   []string `json:"IP_ADDRESSES"`
	Id             string   `json:"ID"`
	Mtu            string   `json:"MTU"`
	Broadcast      string   `json:"BROADCAST"`
	Ipv6_addresses []string `json:"IPV6_ADDRESSES"`
	Interface      string   `json:"INTERFACE"`
	State          string   `json:"STATE"`
	Qlen           string   `json:"QLEN"`
	Qdisc          string   `json:"QDISC"`
}

var IpAddressShow_Template string = `Value Required ID (\d+)
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
