package cisco_xr

type ShowMplsLdpNeighborBrief struct {
	Gr             string `json:"GR"`
	Labels_ipv4    string `json:"LABELS_IPV4"`
	Addresses_ipv6 string `json:"ADDRESSES_IPV6"`
	Labels_ipv6    string `json:"LABELS_IPV6"`
	Peer           string `json:"PEER"`
	Nsr            string `json:"NSR"`
	Uptime         string `json:"UPTIME"`
	Discovery_ipv4 string `json:"DISCOVERY_IPV4"`
	Discovery_ipv6 string `json:"DISCOVERY_IPV6"`
	Addresses_ipv4 string `json:"ADDRESSES_IPV4"`
}

var ShowMplsLdpNeighborBrief_Template string = `Value PEER (\S+)
Value GR (\w)
Value NSR (\S+)
Value UPTIME (\S+)
Value DISCOVERY_IPV4 (\d+)
Value DISCOVERY_IPV6 (\d+)
Value ADDRESSES_IPV4 (\d+)
Value ADDRESSES_IPV6 (\d+)
Value LABELS_IPV4 (\d+)
Value LABELS_IPV6 (\d+)

Start
  ^${PEER}\s+${GR}\s+${NSR}\s+${UPTIME}\s+${DISCOVERY_IPV4}\s+${DISCOVERY_IPV6}\s+${ADDRESSES_IPV4}\s+${ADDRESSES_IPV6}\s+${LABELS_IPV4}\s+${LABELS_IPV6}\s*$$ -> Record
`
