package allied_telesis

type AwplusShowLldpNeighborsDetail struct {
	Management_ip             string `json:"MANAGEMENT_IP"`
	Local_interface           string `json:"LOCAL_INTERFACE"`
	Chassis_id                string `json:"CHASSIS_ID"`
	Neighbor_port_description string `json:"NEIGHBOR_PORT_DESCRIPTION"`
	Neighbor_port_type        string `json:"NEIGHBOR_PORT_TYPE"`
	Neighbor_port_id          string `json:"NEIGHBOR_PORT_ID"`
	Neighbor                  string `json:"NEIGHBOR"`
}

var AwplusShowLldpNeighborsDetail_Template string = `Value LOCAL_INTERFACE (\w+\.\w+\.\w+)
Value CHASSIS_ID (.+)
Value NEIGHBOR_PORT_DESCRIPTION (.+)
Value NEIGHBOR_PORT_TYPE (.+)
Value NEIGHBOR_PORT_ID (.+)
Value NEIGHBOR (.+)
Value MANAGEMENT_IP (.+)

Start
  ^LLDP\s+Detailed\s+Neighbor\s+Information -> Neighbor
  ^. -> Error

Neighbor
  ^Local -> Continue.Record
  ^Local\s+${LOCAL_INTERFACE}:
  ^\s+Neighbors\s+table\s+last\s+updated
  ^\s+TTL
  ^\s+System\s+Description
  ^\s+System\s+Capabilities
  ^\s+Port\s+VLAN\s+ID
  ^\s+Port\s+&\s+Protocol\s+VLAN
  ^\s+Protocol\s+IDs
  ^\s+MAC/PHY\s+Auto-negotiation
  ^\s+Advertised\s+Capability
  ^\s+Operational\s+MAU\s+Type
  ^\s+Power\s+Via\s+MDI
  ^\s+Link\s+Aggregation
  ^\s+Maximum\s+Frame\s+Size
  ^\s+LLDP-MED\s+Device\s+Type
  ^\s+LLDP-MED\s+Capabilities
  ^\s+Network\s+Policy
  ^\s+Location\s+Identification
  ^\s+Extended\s+Power\s+Via\s+MDI
  ^\s+Firmware\s+Revision
  ^\s+Software\s+Revision
  ^\s+Serial\s+Number
  ^\s+Manufacturer\s+Name
  ^\s+Model\s+Name
  ^\s+Asset\s+ID
  ^\s+Port\s+Class
  ^\s+Pair\s+Control\s+Ability
  ^\s+Power\s+Class
  ^\s+Aggregated\s+Port-ID
  ^\s+Chassis\s+ID\s+\S+\s+${CHASSIS_ID}
  ^\s+Port\s+ID\s+Type\s+\.*\s+${NEIGHBOR_PORT_TYPE}
  ^\s+Port\s+ID\s+\.*\s+${NEIGHBOR_PORT_ID}
  ^\s+Port\s+Description\s+\.*\s+${NEIGHBOR_PORT_DESCRIPTION}
  ^\s+System\s+Name\s+\S+\s+${NEIGHBOR}
  ^\s+Management\s+Addresses\s+\S+\s+${MANAGEMENT_IP}
  ^\s+VLAN\s+Names
  ^\s+Inventory\s+Management
  ^\s+Power\s+Source
  ^\s+Power\s+Priority
  ^\s+Power\s+Value
  ^\s+Hardware\s+Revision
  ^\s*------
  # Drop multi line options
  ^\s{8,}
  ^. -> Error
`
