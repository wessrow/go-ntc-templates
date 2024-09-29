package oneaccess_oneos 

type ShowRunningConfigIpRoute struct {
	Network	string	`json:"NETWORK"`
	Mask	string	`json:"MASK"`
	Nexthop_ip	string	`json:"NEXTHOP_IP"`
	Nexthop_interface	string	`json:"NEXTHOP_INTERFACE"`
	Vrf	string	`json:"VRF"`
}

var ShowRunningConfigIpRoute_Template = `Value Required NETWORK (\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})
Value MASK (\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})
Value NEXTHOP_IP (\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3})
Value NEXTHOP_INTERFACE (\w[\w\-\.:\/]+( [\d\.]+)?)
Value VRF (\S+)

Start
  ^ip\sroute(\svrf\s${VRF})?\s${NETWORK}\s${MASK}\s${NEXTHOP_IP} -> Record
  ^ip\sroute(\svrf\s${VRF})?\s${NETWORK}\s${MASK}\s${NEXTHOP_INTERFACE} -> Record
  ^\s*$$
  ^. -> Error
`