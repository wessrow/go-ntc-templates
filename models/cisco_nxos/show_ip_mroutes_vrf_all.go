package cisco_nxos

type ShowIpMroutesVrfAll struct {
	Incoming_interface string   `json:"INCOMING_INTERFACE"`
	Rpf_neighbor       string   `json:"RPF_NEIGHBOR"`
	Last_uptime        string   `json:"LAST_UPTIME"`
	Group              string   `json:"GROUP"`
	Source             string   `json:"SOURCE"`
	Vrf                string   `json:"VRF"`
	Oif_list           []string `json:"OIF_LIST"`
}

var ShowIpMroutesVrfAll_Template string = `Value Required GROUP (\S+)
Value SOURCE (\*|\S+)
Value Filldown VRF (\S+)
Value List OIF_LIST (\S+)
Value INCOMING_INTERFACE (\S+)
Value RPF_NEIGHBOR (\S+)
Value LAST_UPTIME (\S+)

Start
  ^\s*$$
  ^IP\s+Multicast\s+Routing\s+Table\s+for\s+VRF\s+"${VRF}"
  ^\( -> Continue.Record
  ^\(${SOURCE},\s+${GROUP}\),\s+uptime:\s+${LAST_UPTIME},.*
  ^\s+Incoming\s+interface:\s+${INCOMING_INTERFACE},\s+RPF\s+nbr:\s+${RPF_NEIGHBOR}
  ^\s+Outgoing\s+interface\s+list:\s+\(count:\s+\d+\)
  ^\s+${OIF_LIST},\s+uptime:\s+${LAST_UPTIME},\s+\w+
  ^. -> Error

# nxos-2# show ip mroute vrf all
# IP Multicast Routing Table for VRF "default"
#
# (10.0.13.10/32, 225.0.0.1/32), uptime: 00:11:35, pim mrib ip 
#   Incoming interface: Ethernet1/2, RPF nbr: 10.0.2.2, internal
#   Outgoing interface list: (count: 0)
#
#
# (*, 232.0.0.0/8), uptime: 23:48:01, pim ip 
#   Incoming interface: Null, RPF nbr: 0.0.0.0
#   Outgoing interface list: (count: 0)
`
