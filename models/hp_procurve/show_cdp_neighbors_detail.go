package hp_procurve

type ShowCdpNeighborsDetail struct {
	Platform             string `json:"PLATFORM"`
	Capabilities         string `json:"CAPABILITIES"`
	Neighbor_interface   string `json:"NEIGHBOR_INTERFACE"`
	Neighbor_description string `json:"NEIGHBOR_DESCRIPTION"`
	Local_interface      string `json:"LOCAL_INTERFACE"`
	Chassis_id           string `json:"CHASSIS_ID"`
	Mgmt_address         string `json:"MGMT_ADDRESS"`
}

var ShowCdpNeighborsDetail_Template string = `Value Required LOCAL_INTERFACE (.+?)
Value CHASSIS_ID (.+?)
Value MGMT_ADDRESS (.+?)
Value PLATFORM (.+?)
Value CAPABILITIES (.+?)
Value NEIGHBOR_INTERFACE (.+?)
Value NEIGHBOR_DESCRIPTION (.+?)

Start
  ^\s*CDP\s+neighbors\s+information
  ^\s*Port\s*:\s*${LOCAL_INTERFACE}\s*$$
  ^\s*Device\s+ID\s*:\s*${CHASSIS_ID}\s*$$
  ^\s*Address\s+Type
  ^\s*Address\s*:\s*${MGMT_ADDRESS}\s*$$
  ^\s*Platform\s*:\s*${PLATFORM}\s*$$
  ^\s*Capability\s*:\s*${CAPABILITIES}\s*$$
  ^\s*Device\s+Port\s*:\s*${NEIGHBOR_INTERFACE}\s*$$
  ^\s*Version\s*:\s*${NEIGHBOR_DESCRIPTION}\s*$$ -> Record
  ^\s*$$
  ^\s*----
  ^. -> Error
`
