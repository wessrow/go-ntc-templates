package hp_comware 

type DisplayIpRoutingTable struct {
	Network	string	`json:"NETWORK"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Protocal	string	`json:"PROTOCAL"`
	Distance	string	`json:"DISTANCE"`
	Metric	string	`json:"METRIC"`
	Nexthop_ip	string	`json:"NEXTHOP_IP"`
	Nexthop_if	string	`json:"NEXTHOP_IF"`
}

var DisplayIpRoutingTable_Template = `Value Filldown NETWORK (\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})
Value Filldown PREFIX_LENGTH (\d{1,2})
Value Filldown PROTOCAL (\w+)
Value Filldown DISTANCE (\d{1,3})
Value Filldown METRIC (\d+)
Value Required NEXTHOP_IP (\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})
Value NEXTHOP_IF (\S+)

Start
  ^Destination\/Mask -> Routes

Routes
  ^${NETWORK}\/${PREFIX_LENGTH}\s+${PROTOCAL}\s+${DISTANCE}\s+${METRIC}\s+${NEXTHOP_IP}\s+${NEXTHOP_IF} -> Record
  # for comware v7 ecmp
  ^\s+${NEXTHOP_IP}\s+${NEXTHOP_IF} -> Record
  # for comware v5 ecmp
  ^\s+${PROTOCAL}\s+${DISTANCE}\s+${METRIC}\s+${NEXTHOP_IP}\s+${NEXTHOP_IF} -> Record
  ^\s*$$
  ^. -> Error
`