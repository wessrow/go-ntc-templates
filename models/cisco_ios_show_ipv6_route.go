package models

type CiscoIosShowIpv6Route struct {
	Protocol	string	`json:"PROTOCOL"`
	Network	string	`json:"NETWORK"`
	Distance	string	`json:"DISTANCE"`
	Metric	string	`json:"METRIC"`
	Nexthop_ip	string	`json:"NEXTHOP_IP"`
	Nexthop_if	string	`json:"NEXTHOP_IF"`
}

var CiscoIosShowIpv6Route_Template = `Value Filldown PROTOCOL (\w{1,4})
Value Filldown NETWORK ([A-Za-z0-9:]+/\d{1,3}(?:/\d{2})?)
Value Filldown DISTANCE (\d+)
Value Filldown METRIC (\d+)
Value NEXTHOP_IP ([A-F0-9:]+)
Value NEXTHOP_IF ([\w\d\s/]+)


Start
  # Match regular ipv6 routes
  ^IPv6[\d\w\-\s,:]\s*
  ^Codes:[\d\w\-\s,]\s*
  ^\s{7}[\d\w\-\s,]\s*
  ^${PROTOCOL}\s+${NETWORK}\s+(\[${DISTANCE}/${METRIC}])\s*$$
  ^\s*via\s+${NEXTHOP_IP}(,\s+${NEXTHOP_IF}.*)?\s*$$ -> Record
  ^\s*via\s+${NEXTHOP_IF}.*$$ -> Record
  ^\s* -> Clearall
  ^.*$$ -> Error

EOF
`