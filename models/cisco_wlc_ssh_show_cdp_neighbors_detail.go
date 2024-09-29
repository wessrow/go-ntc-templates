package models

type CiscoWlcSshShowCdpNeighborsDetail struct {
	Chassis_id	string	`json:"CHASSIS_ID"`
	Mgmt_address	string	`json:"MGMT_ADDRESS"`
	Platform	string	`json:"PLATFORM"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_description	string	`json:"NEIGHBOR_DESCRIPTION"`
	Capabilities	string	`json:"CAPABILITIES"`
}

var CiscoWlcSshShowCdpNeighborsDetail_Template = `Value Required CHASSIS_ID (\S+)
Value MGMT_ADDRESS (\d+\.\d+\.\d+\.\d+|\w+\.\w+\.\w+)
Value PLATFORM (.*)
Value NEIGHBOR_INTERFACE (.*)
Value LOCAL_INTERFACE (.*)
Value NEIGHBOR_DESCRIPTION (.*$)
Value CAPABILITIES (.*)

Start
  ^Device ID: ${CHASSIS_ID}
  ^Entry address\(es\)\s*:\s* -> ParseIP
  ^Platform\s*:\s*${PLATFORM}\s*,\s*Capabilities\s*:\s*${CAPABILITIES}
  ^Interface: ${LOCAL_INTERFACE},  Port ID \(outgoing port\): ${NEIGHBOR_INTERFACE}
  ^Version : -> GetVersion

ParseIP
  ^.*IP address: ${MGMT_ADDRESS} -> Start
  ^Platform\s*:\s*${PLATFORM}\s*,\s*Capabilities\s*:\s*${CAPABILITIES} -> Start
  ^.* -> Start

GetVersion
  ^${NEIGHBOR_DESCRIPTION} -> Record Start

`