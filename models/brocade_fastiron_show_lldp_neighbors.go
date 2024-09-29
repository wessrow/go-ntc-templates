package models

type BrocadeFastironShowLldpNeighbors struct {
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Holdtime	string	`json:"HOLDTIME"`
	Capabilities	string	`json:"CAPABILITIES"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
}

var BrocadeFastironShowLldpNeighbors_Template = `Value Required NEIGHBOR_NAME (\S+)
Value Required LOCAL_INTERFACE (\S+)
Value Required HOLDTIME (\d+)
Value Required CAPABILITIES (\S+)
Value Required NEIGHBOR_INTERFACE (\S+) 

Start
  ^Capability.*
  ^\s+\(.*
  ^Device\s+ID\s+Local\s+Intf\s+Hold-time\s+Capability\s+Port\s+ID\s*$$
  ^${NEIGHBOR_NAME}\s+${LOCAL_INTERFACE}\s+${HOLDTIME}\s+${CAPABILITIES}\s+${NEIGHBOR_INTERFACE} -> Record
  ^Total.*
  ^\s*$$
  ^. -> Error

`