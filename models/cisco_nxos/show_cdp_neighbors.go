package cisco_nxos 

type ShowCdpNeighbors struct {
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Holdtime	string	`json:"HOLDTIME"`
	Capabilities	string	`json:"CAPABILITIES"`
	Platform	string	`json:"PLATFORM"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
}

var ShowCdpNeighbors_Template = `Value Required,Filldown NEIGHBOR_NAME (\S+)
Value Required LOCAL_INTERFACE (\S+)
Value HOLDTIME (\d+)
Value CAPABILITIES ([\S+\s+]+?)
Value PLATFORM (\S+)
Value NEIGHBOR_INTERFACE (\S+)

Start
  ^Device.*ID -> CDP

CDP
  ^${NEIGHBOR_NAME}\s+${LOCAL_INTERFACE}\s+${HOLDTIME}\s+${CAPABILITIES}\s+${PLATFORM}\s+${NEIGHBOR_INTERFACE}\s*$$ -> Record
  ^${NEIGHBOR_NAME}\s*$$
  ^\s+${LOCAL_INTERFACE}\s+${HOLDTIME}\s+${CAPABILITIES}\s+${PLATFORM}\s+${NEIGHBOR_INTERFACE}\s*$$ -> Record


`