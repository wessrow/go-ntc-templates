package cisco_xr

type ShowCdpNeighborsDetail struct {
	Local_interface      string `json:"LOCAL_INTERFACE"`
	Neighbor_description string `json:"NEIGHBOR_DESCRIPTION"`
	Capabilities         string `json:"CAPABILITIES"`
	Chassis_id           string `json:"CHASSIS_ID"`
	Neighbor_name        string `json:"NEIGHBOR_NAME"`
	Mgmt_address         string `json:"MGMT_ADDRESS"`
	Platform             string `json:"PLATFORM"`
	Neighbor_interface   string `json:"NEIGHBOR_INTERFACE"`
}

var ShowCdpNeighborsDetail_Template string = `Value Required CHASSIS_ID (\S+)
Value NEIGHBOR_NAME (.*)
Value MGMT_ADDRESS (.*)
Value PLATFORM (.*)
Value NEIGHBOR_INTERFACE (.*)
Value LOCAL_INTERFACE (.*)
Value NEIGHBOR_DESCRIPTION (.*)
Value CAPABILITIES (.*)

Start
  ^Device ID: ${CHASSIS_ID}
  ^SysName : ${NEIGHBOR_NAME}
  ^Entry address\(es\): -> GetIP
  ^Platform: ${PLATFORM},  Capabilities: ${CAPABILITIES}
  ^Interface: ${LOCAL_INTERFACE}
  ^Port ID \(outgoing port\): ${NEIGHBOR_INTERFACE}
  ^Version : -> GetVersion

GetIP
  ^.*IP.+address: ${MGMT_ADDRESS} -> Start

GetVersion
  ^${NEIGHBOR_DESCRIPTION} -> Record Start

`
