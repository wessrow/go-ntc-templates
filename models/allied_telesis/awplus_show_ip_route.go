package allied_telesis

type AwplusShowIpRoute struct {
	Type        string `json:"TYPE"`
	Distance    string `json:"DISTANCE"`
	Metric      string `json:"METRIC"`
	Nexthop_vrf string `json:"NEXTHOP_VRF"`
	Uptime      string `json:"UPTIME"`
	Protocol    string `json:"PROTOCOL"`
	Network     string `json:"NETWORK"`
	Mask        string `json:"MASK"`
	Nexthop_ip  string `json:"NEXTHOP_IP"`
	Nexthop_if  string `json:"NEXTHOP_IF"`
	Vrf         string `json:"VRF"`
}

var AwplusShowIpRoute_Template string = `Value Filldown VRF (\S+)
Value PROTOCOL (\w)
Value TYPE (\w{0,2})
Value Required NETWORK (\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})
Value Required MASK (\d{1,2})
Value DISTANCE (\d+)
Value METRIC (\d+)
Value NEXTHOP_IP (\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})
Value NEXTHOP_VRF (\S+)
Value NEXTHOP_IF ([A-Za-z][\w\-\.:/]+)
Value UPTIME (\d[\w:\.]+)

Start
  ^Codes:
  ^\s{4,}O\s-\sOSPF
  ^\s{4,}N1\s-\sOSPF
  ^\s{4,}E1\s-\sOSPF
  ^\s{4,}i\s-\sIS-IS
  ^\s{4,}\*\s-\scandidate -> Routes
  ^. -> Error

Routes
  ^\[VRF:\s${VRF}\]$$
  ^\s*Gateway
  # S*      0.0.0.0/0 [1/0] via 192.168.254.251, vlan954
  ^${PROTOCOL}(\s|\*)${TYPE}\s+${NETWORK}\/${MASK}\s\[${DISTANCE}/${METRIC}\]\svia\s${NEXTHOP_IP},\s${NEXTHOP_IF} -> Record
  # C       10.255.255.0/24 is directly connected, vlan795
  ^${PROTOCOL}(\s|\*)${TYPE}\s+${NETWORK}\/${MASK}\sis\sdirectly\sconnected,\s${NEXTHOP_IF} -> Record
  ^. -> Error
`
