package models

type CiscoNxosShowLldpNeighbors struct {
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Capabilities	string	`json:"CAPABILITIES"`
}

var CiscoNxosShowLldpNeighbors_Template = `Value NEIGHBOR_NAME (\S+)
Value LOCAL_INTERFACE (\S+)
Value NEIGHBOR_INTERFACE (\S+)
Value CAPABILITIES (\w+)

Start
  ^Device.*ID -> LLDP

LLDP
  ^${NEIGHBOR_NAME}$$
  ^${NEIGHBOR_NAME}\s+${LOCAL_INTERFACE}\s+\d+\s+(${CAPABILITIES}?\s)+\s+${NEIGHBOR_INTERFACE} -> Record
  ^\s+${LOCAL_INTERFACE}\s+\d+\s+(${CAPABILITIES}?\s)+\s+${NEIGHBOR_INTERFACE} -> Record

`