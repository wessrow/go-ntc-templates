package arista_eos 

type ShowLldpNeighbors struct {
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
}

var ShowLldpNeighbors_Template = `Value NEIGHBOR_NAME (\S+)
Value LOCAL_INTERFACE (\S+)
Value NEIGHBOR_INTERFACE (\S+)

Start
  ^Port.*TTL -> LLDP

LLDP
  # Skip the hyphen header line
  ^--------.*$$
  ^${LOCAL_INTERFACE}\s+${NEIGHBOR_NAME}\s+${NEIGHBOR_INTERFACE}\s+.* -> Record
  ^\s*$$
  ^. -> Error
`