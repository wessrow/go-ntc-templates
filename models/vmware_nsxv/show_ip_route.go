package vmware_nsxv

type ShowIpRoute struct {
	Ip_address    string `json:"IP_ADDRESS"`
	Prefix_length string `json:"PREFIX_LENGTH"`
	Distance      string `json:"DISTANCE"`
	Metric        string `json:"METRIC"`
	Nexthop       string `json:"NEXTHOP"`
	Protocol      string `json:"PROTOCOL"`
	Type          string `json:"TYPE"`
}

var ShowIpRoute_Template string = `Value PROTOCOL ([OiBCS])
Value TYPE (\w{0,2})
Value IP_ADDRESS (\d+(\.\d+){3})
Value PREFIX_LENGTH (\d{1,2})
Value DISTANCE (\d+)
Value METRIC (\d+)
Value NEXTHOP (\d+(\.\d+){3})

Start
	^${PROTOCOL}\s+${TYPE}\s+${IP_ADDRESS}\/${PREFIX_LENGTH}\s+\[${DISTANCE}\/${METRIC}\]\s+via\s+${NEXTHOP} -> Next.Record
`
