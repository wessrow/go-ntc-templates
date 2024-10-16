package cisco_s300

type ShowLldpNeighbors struct {
	Neighbor_name      []string `json:"NEIGHBOR_NAME"`
	Local_interface    string   `json:"LOCAL_INTERFACE"`
	Neighbor_interface []string `json:"NEIGHBOR_INTERFACE"`
	Chassis_id         string   `json:"CHASSIS_ID"`
}

var ShowLldpNeighbors_Template string = `Value List NEIGHBOR_NAME (\S+)
Value LOCAL_INTERFACE (\S+)
Value List NEIGHBOR_INTERFACE (\S+)
Value CHASSIS_ID (([0-9a-f]{2}[:-]){5}([0-9a-f]{2}))

Start
  ^System\s+capability\s+legend
  ^\S+\s*-
  ^\s*Port\s+Device\s+ID\s+Port\s+ID\s+System\s+Name\s+Capabilities\s+TTL -> Begin
  ^\s*$$
  ^. -> Error

Begin
  ^\S+ -> Continue.Record
  ^${LOCAL_INTERFACE}\s+${CHASSIS_ID}\s+${NEIGHBOR_INTERFACE}\s+(${NEIGHBOR_NAME})?\s+\S+\s+\d+
  ^\s{28}${NEIGHBOR_INTERFACE}\s*$$
  ^\s{42}${NEIGHBOR_NAME}\s*$$
  ^\s{28}${NEIGHBOR_INTERFACE}\s+${NEIGHBOR_NAME}\s*$$
  ^-+
  ^\s*$$
  ^. -> Error
`
