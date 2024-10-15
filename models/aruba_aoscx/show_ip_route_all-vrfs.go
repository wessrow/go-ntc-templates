package aruba_aoscx

type ShowIpRouteAllVrfs struct {
	Status        []string `json:"STATUS"`
	Ip_address    string   `json:"IP_ADDRESS"`
	Prefix_length string   `json:"PREFIX_LENGTH"`
	Vrf           string   `json:"VRF"`
	Interface     []string `json:"INTERFACE"`
	Metric        []string `json:"METRIC"`
}

var ShowIpRouteAllVrfs_Template string = `Value Filldown IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value Filldown PREFIX_LENGTH (\d+)
Value Filldown VRF (\S+)
Value List INTERFACE (\S+)
Value List METRIC (\[\S+\])
Value List STATUS (\w+)

Start
  ^\s*Displaying.*
  ^\s*$$
  ^\s*\S+\s+denotes.*
  ^\d+\.\d+\.\d+\.\d+\/\d+\W\s+vrf\s+\S+ -> Continue.Record
  ^${IP_ADDRESS}\/${PREFIX_LENGTH}\W\s+vrf\s+${VRF}
  ^\s+via\s+${INTERFACE},\s+${METRIC},\s+${STATUS}
  ^. -> Error
`
