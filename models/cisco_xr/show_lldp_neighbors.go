package cisco_xr 

type ShowLldpNeighbors struct {
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
}

var ShowLldpNeighbors_Template = `Value Required NEIGHBOR_NAME (\S+)
Value Required LOCAL_INTERFACE (\S+)
Value Required NEIGHBOR_INTERFACE (\S+)

Start
  ^Device.*ID -> LLDP

LLDP
  ^${NEIGHBOR_NAME}\s+${LOCAL_INTERFACE}\s+\d+\s+(.+?)\s+${NEIGHBOR_INTERFACE}$$ -> Record
  ^${NEIGHBOR_NAME}
  ^\s+${LOCAL_INTERFACE}\s+\d+\s+(.*?)\s+${NEIGHBOR_INTERFACE} -> Record
  ^Total entries
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"
`