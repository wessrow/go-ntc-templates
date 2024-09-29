package cisco_ios 

type ShowCdpNeighborsDetail struct {
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Mgmt_address	string	`json:"MGMT_ADDRESS"`
	Platform	string	`json:"PLATFORM"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_description	string	`json:"NEIGHBOR_DESCRIPTION"`
	Capabilities	string	`json:"CAPABILITIES"`
}

var ShowCdpNeighborsDetail_Template = `Value Required NEIGHBOR_NAME (\S+)
Value MGMT_ADDRESS (\d+\.\d+\.\d+\.\d+|\w+\.\w+\.\w+)
Value PLATFORM (.*)
Value NEIGHBOR_INTERFACE (.*)
Value LOCAL_INTERFACE (.*)
Value NEIGHBOR_DESCRIPTION (.*$)
Value CAPABILITIES (.+?)

Start
  ^Device ID: ${NEIGHBOR_NAME}
  ^Entry address\(es\)\s*:\s* -> ParseIP
  ^Platform\s*:\s*${PLATFORM}\s*,\s*Capabilities\s*:\s*${CAPABILITIES}\s+$$
  ^Platform\s*:\s*${PLATFORM}\s*,\s*Capabilities\s*:\s*${CAPABILITIES}$$
  ^Interface: ${LOCAL_INTERFACE},  Port ID \(outgoing port\): ${NEIGHBOR_INTERFACE}
  ^Version : -> GetVersion
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

ParseIP
  ^.*IP address: ${MGMT_ADDRESS} -> Start
  ^Platform\s*:\s*${PLATFORM}\s*,\s*Capabilities\s*:\s*${CAPABILITIES}\s+$$ -> Start
  ^Platform\s*:\s*${PLATFORM}\s*,\s*Capabilities\s*:\s*${CAPABILITIES}$$ -> Start
  ^.* -> Start

GetVersion
  ^${NEIGHBOR_DESCRIPTION} -> Record Start
`