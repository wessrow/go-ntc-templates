package broadcom_icos

type ShowIsdpNeighbors struct {
	Capabilities       string `json:"CAPABILITIES"`
	Platform           string `json:"PLATFORM"`
	Neighbor_interface string `json:"NEIGHBOR_INTERFACE"`
	Neighbor_name      string `json:"NEIGHBOR_NAME"`
	Local_interface    string `json:"LOCAL_INTERFACE"`
	Holdtime           string `json:"HOLDTIME"`
}

var ShowIsdpNeighbors_Template string = `Value NEIGHBOR_NAME (\S+)
Value LOCAL_INTERFACE (\S+)
Value HOLDTIME (\S+)
Value CAPABILITIES ([\w]{1}(?:\s[\w]){0,2})
Value PLATFORM ((\S+\s\S+)|(\S+))
Value NEIGHBOR_INTERFACE (.+)

Start
  # Captures show isdp neighbors for:
  # Accton AS4610-54P, Accton AS5610-52X, Quanta LY2R, Quanta LB9, DNI AG3448P-R   
  # Raw data is the same in the case of all those devices
  ^Capability\sCodes:
  ^\s+\S+\s+-
  ^\s*Device\s+ID\s+Intf\s+Holdtime\s+Capability\s+Platform\s+Port\s+ID$$
  ^-+
  ^\s*${NEIGHBOR_NAME}\s+${LOCAL_INTERFACE}\s+${HOLDTIME}\s+${CAPABILITIES}\s+${PLATFORM}\s+${NEIGHBOR_INTERFACE} -> Record
  ^\s*$$
  ^. -> Error
`
