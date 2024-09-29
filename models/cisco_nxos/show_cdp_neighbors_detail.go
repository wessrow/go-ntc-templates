package cisco_nxos 

type ShowCdpNeighborsDetail struct {
	Chassis_id	string	`json:"CHASSIS_ID"`
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Mgmt_address	string	`json:"MGMT_ADDRESS"`
	Platform	string	`json:"PLATFORM"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_description	string	`json:"NEIGHBOR_DESCRIPTION"`
	Interface_ip	string	`json:"INTERFACE_IP"`
	Capabilities	string	`json:"CAPABILITIES"`
}

var ShowCdpNeighborsDetail_Template = `Value Required CHASSIS_ID (.*)
Value NEIGHBOR_NAME (.*)
Value MGMT_ADDRESS (.*)
Value PLATFORM (.*)
Value NEIGHBOR_INTERFACE (.*)
Value LOCAL_INTERFACE (.*)
Value NEIGHBOR_DESCRIPTION (.*)
Value INTERFACE_IP (.*)
Value CAPABILITIES (.*[^\s])

Start
  ^Device ID -> Continue.Record
  ^Device ID:${CHASSIS_ID}
  ^System Name: ${NEIGHBOR_NAME}
  ^Interface address\(es\):\s*(^[1-9]\d*|$$) -> GetInterfaceIP
  ^Mgmt address\(es\): -> GetIP
  ^Platform: ${PLATFORM}, Capabilities: ${CAPABILITIES}\s*$$
  ^Interface: ${LOCAL_INTERFACE}, Port ID \(outgoing port\): ${NEIGHBOR_INTERFACE}
  ^Version: -> GetVersion

GetIP
  ^.*IP.+Address: ${MGMT_ADDRESS} -> Start

GetInterfaceIP
  ^.*IP.+Address: ${INTERFACE_IP} -> Start

GetVersion
  ^${NEIGHBOR_DESCRIPTION} -> Start
`