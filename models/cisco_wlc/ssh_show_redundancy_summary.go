package cisco_wlc

type SshShowRedundancySummary struct {
	Local_state                 string `json:"LOCAL_STATE"`
	Peer_state                  string `json:"PEER_STATE"`
	Unit_id                     string `json:"UNIT_ID"`
	Redundancy_port             string `json:"REDUNDANCY_PORT"`
	Avg_redundancy_peer_latency string `json:"AVG_REDUNDANCY_PEER_LATENCY"`
	Redundancy_mode             string `json:"REDUNDANCY_MODE"`
	Redundancy_state            string `json:"REDUNDANCY_STATE"`
	Mobility_mac                string `json:"MOBILITY_MAC"`
	Bulksync_status             string `json:"BULKSYNC_STATUS"`
	Avg_mgmt_gw_latency         string `json:"AVG_MGMT_GW_LATENCY"`
	Unit                        string `json:"UNIT"`
}

var SshShowRedundancySummary_Template string = `Value REDUNDANCY_MODE (\S+\s+\S+)
Value LOCAL_STATE (\S+)
Value PEER_STATE (\S+\s+\S+)
Value UNIT (\S+)
Value UNIT_ID (([\da-fA-F]{2}\:?){6})
Value REDUNDANCY_STATE (\S+)
Value MOBILITY_MAC (([\da-fA-F]{2}\:?){6})
Value REDUNDANCY_PORT (\S+)
Value BULKSYNC_STATUS (\S+)
Value AVG_REDUNDANCY_PEER_LATENCY (\d+)
Value AVG_MGMT_GW_LATENCY (\d+)

Start
  ^\s*Redundancy Mode\s+=\s+${REDUNDANCY_MODE}\s*$$
  ^\s*Local State\s+=\s+${LOCAL_STATE}\s*$$
  ^\s*Peer State\s+=\s+${PEER_STATE}\s*$$
  ^\s*Unit\s+=\s+${UNIT}\s*$$
  ^\s*Unit ID\s+=\s+${UNIT_ID}\s*$$
  ^\s*Redundancy State\s+=\s+${REDUNDANCY_STATE}\s*$$
  ^\s*Mobility MAC\s+=\s+${MOBILITY_MAC}\s*$$
  ^\s*Redundancy Port\s+=\s+${REDUNDANCY_PORT}\s*$$
  ^\s*BulkSync Status\s+=\s+${BULKSYNC_STATUS}\s*$$
  ^\s*Average Redundancy Peer Reachability Latency\s+=\s+${AVG_REDUNDANCY_PEER_LATENCY}\s+\S+\s+Seconds\s*$$
  ^\s*Average Management Gateway Reachability Latency\s+=\s+${AVG_MGMT_GW_LATENCY}\s+\S+\s+Seconds\s*$$
`
