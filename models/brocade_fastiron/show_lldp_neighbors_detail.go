package brocade_fastiron 

type ShowLldpNeighborsDetail struct {
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_port_id	string	`json:"NEIGHBOR_PORT_ID"`
	Holdtime	string	`json:"HOLDTIME"`
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Vlan_id	string	`json:"VLAN_ID"`
	Mgmt_address	string	`json:"MGMT_ADDRESS"`
	Neighbor_lacp_index	string	`json:"NEIGHBOR_LACP_INDEX"`
	Neighbor_max_frame_size	string	`json:"NEIGHBOR_MAX_FRAME_SIZE"`
	Neighbor_mau	string	`json:"NEIGHBOR_MAU"`
	Capabilities	string	`json:"CAPABILITIES"`
}

var ShowLldpNeighborsDetail_Template = `Value Filldown LOCAL_INTERFACE ([0-9\/]+)
Value Required NEIGHBOR_PORT_ID (\w+\.\w+\.\w+)
Value HOLDTIME (\d+)
Value NEIGHBOR_NAME (.+)
Value NEIGHBOR_INTERFACE ([a-zA-Z0-9\/\s]+)
Value VLAN_ID (\d+|none)
Value MGMT_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value NEIGHBOR_LACP_INDEX (\d+)
Value NEIGHBOR_MAX_FRAME_SIZE (\d+)
Value NEIGHBOR_MAU ([a-zA-Z0-9\-]+)
Value CAPABILITIES (.+)

Start
  ^\w -> Continue.Record
  ^Local\s+port:\s+${LOCAL_INTERFACE}
  ^\s+\+\s+System\s+name\s+:\s+"${NEIGHBOR_NAME}"
  ^\s+\+\s+Port\s+description\s+:\s+"${NEIGHBOR_INTERFACE}"
  ^\s+Enabled\s+capabilities\s*:\s+${CAPABILITIES}
  ^\s+\+\s+Port\s+VLAN\s+ID:\s+${VLAN_ID}
  ^\s+\+\s+Management\s+address\s+\(IPv4\):\s+${MGMT_ADDRESS}
  ^\s+\+\s+Link\s+aggregation:\s+aggregated\s+\(aggregated\s+port\s+ifIndex:\s+${NEIGHBOR_LACP_INDEX}\)
  ^\s+\+\s+Maximum\s+frame\s+size:\s+${NEIGHBOR_MAX_FRAME_SIZE}
  ^\s+Operational\s+MAU\s+type\s+:\s+${NEIGHBOR_MAU}
  ^\s+Neighbor:\s+${NEIGHBOR_PORT_ID},\s+TTL\s+${HOLDTIME}
`