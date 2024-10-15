package cisco_ios

type ShowCdpNeighbors struct {
	Platform           string `json:"PLATFORM"`
	Neighbor_interface string `json:"NEIGHBOR_INTERFACE"`
	Neighbor_name      string `json:"NEIGHBOR_NAME"`
	Local_interface    string `json:"LOCAL_INTERFACE"`
	Capabilities       string `json:"CAPABILITIES"`
}

var ShowCdpNeighbors_Template string = `Value Required NEIGHBOR_NAME (\S+)
Value LOCAL_INTERFACE (\S+(?:\s\S+)?)
Value CAPABILITIES ((?:\w\s)*\w)
Value PLATFORM ((?:[IiPp]{2}\s)?\S+)
Value NEIGHBOR_INTERFACE (.+?)

Start
  ^Device.*ID -> CDP
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

CDP
  ^${NEIGHBOR_NAME}$$
  ^\s*${LOCAL_INTERFACE}\s+\d+\s+${CAPABILITIES}\s+${PLATFORM}\s+${NEIGHBOR_INTERFACE}\s*$$ -> Record
  ^${NEIGHBOR_NAME}\s+${LOCAL_INTERFACE}\s+\d+\s+${CAPABILITIES}\s+${PLATFORM}\s+${NEIGHBOR_INTERFACE}\s*$$ -> Record
`
