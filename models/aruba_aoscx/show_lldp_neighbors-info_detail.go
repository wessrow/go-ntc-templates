package aruba_aoscx

type ShowLldpNeighborsInfoDetail struct {
	Local_interface        string `json:"LOCAL_INTERFACE"`
	Chassis_id             string `json:"CHASSIS_ID"`
	Neighbor_port_id       string `json:"NEIGHBOR_PORT_ID"`
	Neighbor_interface     string `json:"NEIGHBOR_INTERFACE"`
	Mgmt_address           string `json:"MGMT_ADDRESS"`
	Neighbor_name          string `json:"NEIGHBOR_NAME"`
	Neighbor_description   string `json:"NEIGHBOR_DESCRIPTION"`
	Capabilities_supported string `json:"CAPABILITIES_SUPPORTED"`
	Capabilities           string `json:"CAPABILITIES"`
}

var ShowLldpNeighborsInfoDetail_Template string = `Value Required LOCAL_INTERFACE (\S+)
Value Required CHASSIS_ID (\S+)
Value NEIGHBOR_NAME (\S+)
Value NEIGHBOR_DESCRIPTION (.+)
Value CAPABILITIES_SUPPORTED (.+)
Value CAPABILITIES (.+)
Value MGMT_ADDRESS (\S+)
Value NEIGHBOR_PORT_ID (\S+)
Value NEIGHBOR_INTERFACE (\S+)


Start
  ^Port\s*:\s*${LOCAL_INTERFACE}
  ^Neighbor\s*(Chassis|System)-Name\s*:\s*${NEIGHBOR_NAME}
  ^Neighbor\s*(Chassis|System)-Description\s*:\s*${NEIGHBOR_DESCRIPTION}
  ^Neighbor\s*Chassis-ID\s*:\s*${CHASSIS_ID}
  ^Neighbor\s*Management-Address\s*:\s*${MGMT_ADDRESS}
  ^Chassis\s*Capabilities\s*Available\s*:\s*${CAPABILITIES_SUPPORTED}
  ^Chassis\s*Capabilities\s*Enabled\s*:\s*${CAPABILITIES}
  ^Neighbor\s*Port-ID\s*:\s*${NEIGHBOR_PORT_ID}
  ^Neighbor\s*Port-Desc\s*:\s*${NEIGHBOR_INTERFACE} 
  ^-+$$ -> Record
  ^Neighbor\.*
  ^TTL\.*
  ^P[D|SE]\.*
  ^Power\.*
  ^MED\s+capabilities
  ^\s+ -> Next
  ^\S+#\.* -> Next
  ^LLDP\.*
  ^Link\s+aggregation 
  ^Aggregation\s+port
  ^Total\.*
  ^=+
  ^. -> Error

`
