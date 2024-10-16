package cisco_ios

type ShowLldpNeighbors struct {
	Neighbor_name      string `json:"NEIGHBOR_NAME"`
	Local_interface    string `json:"LOCAL_INTERFACE"`
	Capabilities       string `json:"CAPABILITIES"`
	Neighbor_interface string `json:"NEIGHBOR_INTERFACE"`
}

var ShowLldpNeighbors_Template string = `Value Required NEIGHBOR_NAME (.{0,20}(?<! ))
Value Required LOCAL_INTERFACE (\S+)
Value CAPABILITIES (\S*)
Value Required NEIGHBOR_INTERFACE (\S+)

Start
  ^Device.*ID -> LLDP
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

LLDP
  ^${NEIGHBOR_NAME}\s*${LOCAL_INTERFACE}\s+\d+\s+${CAPABILITIES}\s+${NEIGHBOR_INTERFACE} -> Record
  ^${NEIGHBOR_NAME}
  ^\s+${LOCAL_INTERFACE}\s+\d+\s+${CAPABILITIES}\s+${NEIGHBOR_INTERFACE} -> Record
`
