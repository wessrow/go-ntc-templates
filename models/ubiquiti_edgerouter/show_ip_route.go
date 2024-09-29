package ubiquiti_edgerouter 

type ShowIpRoute struct {
	Code	string	`json:"CODE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Distance	string	`json:"DISTANCE"`
	Metric	string	`json:"METRIC"`
	Nexthop_ip	string	`json:"NEXTHOP_IP"`
	Interface	string	`json:"INTERFACE"`
}

var ShowIpRoute_Template = `Value Filldown,Required CODE (\w{1,2})
Value Filldown IP_ADDRESS (\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})
Value Filldown PREFIX_LENGTH (\d{1,2})
Value DISTANCE (\d+)
Value METRIC (\d+)
Value NEXTHOP_IP (\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})
Value INTERFACE (\w+)

Start
  ^.*IP\sRoute\sTable -> IP

IP
  ^\s+\*>\s+\[${DISTANCE}/${METRIC}\]\svia\s${NEXTHOP_IP}\s\(recursive is directly connected,\s${INTERFACE}\)\s* -> Record
  ^${CODE}\s+\*>\s${IP_ADDRESS}/${PREFIX_LENGTH}\sis\sdirectly\sconnected,\s${INTERFACE}\s* -> Record
  ^${CODE}\s+\*>\s${IP_ADDRESS}/${PREFIX_LENGTH}\sis\sdirectly\sconnected\s* -> Record
  ^${CODE}\s+\*>\s${IP_ADDRESS}/${PREFIX_LENGTH}\s\[${DISTANCE}/${METRIC}\]\sis\sdirectly\sconnected,\s${INTERFACE}\s* -> Record
  ^${CODE}\s+\*>\s${IP_ADDRESS}/${PREFIX_LENGTH}\s\[${DISTANCE}/${METRIC}\]\svia\s${NEXTHOP_IP}\s\(recursive is directly connected,\s${INTERFACE}\)\s* -> Record
  ^. -> Error

EOF
`