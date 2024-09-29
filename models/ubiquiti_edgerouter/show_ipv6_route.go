package ubiquiti_edgerouter 

type ShowIpv6Route struct {
	Code	string	`json:"CODE"`
	Ipv6_address	string	`json:"IPV6_ADDRESS"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Distance	string	`json:"DISTANCE"`
	Metric	string	`json:"METRIC"`
	Nexthop_ip	string	`json:"NEXTHOP_IP"`
	Interface	string	`json:"INTERFACE"`
	Uptime	string	`json:"UPTIME"`
}

var ShowIpv6Route_Template = `Value CODE (\w{1,2})
Value IPV6_ADDRESS ([A-Za-z0-9:]+)
Value PREFIX_LENGTH (\d{1,3})
Value DISTANCE (\d+)
Value METRIC (\d+)
Value NEXTHOP_IP ([\w:]+)
Value INTERFACE (\w+)
Value UPTIME (\d[\w:\.]+)

Start
  ^.*IP\sRoute\sTable -> IP

IP
  ^${CODE}\s+${IPV6_ADDRESS}/${PREFIX_LENGTH}\s\[${DISTANCE}/${METRIC}\]\svia\s${NEXTHOP_IP},\s${INTERFACE},\s${UPTIME}\s* -> Record
  ^${CODE}\s+${IPV6_ADDRESS}/${PREFIX_LENGTH}\svia\s${NEXTHOP_IP},\s${INTERFACE},\s${UPTIME}\s* -> Record
  ^\s*$$
  ^. -> Error
`