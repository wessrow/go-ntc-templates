package cisco_xr

type ShowLldpNeighborsDetail struct {
	Neighbor                  string `json:"NEIGHBOR"`
	Neighbor_port_description string `json:"NEIGHBOR_PORT_DESCRIPTION"`
	Neighbor_port_id          string `json:"NEIGHBOR_PORT_ID"`
	System_description        string `json:"SYSTEM_DESCRIPTION"`
	Capabilities              string `json:"CAPABILITIES"`
	Management_ip             string `json:"MANAGEMENT_IP"`
	Management_ipv6           string `json:"MANAGEMENT_IPv6"`
	Mac_address               string `json:"MAC_ADDRESS"`
	Local_interface           string `json:"LOCAL_INTERFACE"`
	Chassis_id                string `json:"CHASSIS_ID"`
}

var ShowLldpNeighborsDetail_Template string = `Value LOCAL_INTERFACE (\S+)
Value CHASSIS_ID ([0-9a-fA-F]{4}\.[0-9a-fA-F]{4}\.[0-9a-fA-F]{4})
Value NEIGHBOR_PORT_DESCRIPTION (.*)
Value NEIGHBOR_PORT_ID (.*)
Value NEIGHBOR (.+)
Value SYSTEM_DESCRIPTION (.*)
Value CAPABILITIES (.*)
Value MANAGEMENT_IP (\S+)
Value MANAGEMENT_IPv6 (\S+)
Value MAC_ADDRESS (([0-9a-fA-F]{2}[:]){5}([0-9a-fA-F]{2}))

Start
  ^$$
  ^\S{3}\s+\S{3}\s+\d{1,2}\s+\d+:\d+:\d+
  ^Capability\s+codes\: -> CAPABILITY_CODES
  ^-------------- -> NEIGHBOR
  ^.* -> Error

NEIGHBOR
  ^$$
  ^Local\s+Interface\:\s+${LOCAL_INTERFACE}
  ^Chassis\s+id\:\s+${CHASSIS_ID}
  ^Port\s+id\:\s+${NEIGHBOR_PORT_ID}
  ^Port\s+Description\:\s+${NEIGHBOR_PORT_DESCRIPTION}
  ^System\s+Name\:\s+${NEIGHBOR}
  ^System\s+Description\: -> DESCRIPTION
  ^Time\s+remaining\:\s+
  ^Hold\s+Time\:\s+
  ^Age\:\s+\d+
  ^System\s+Capabilities\:\s+
  ^Enabled\s+Capabilities\:\s+${CAPABILITIES}
  ^Management\s+Addresses\:
  ^\s*IPv4\s+address\:\s+${MANAGEMENT_IP}
  ^\s*IPv6\s+address\:\s+${MANAGEMENT_IPv6}
  ^Peer\s+MAC\s+Address\:\s+${MAC_ADDRESS}
  ^Total\s+entries\s+displayed\:\s\d+ -> Start
  ^-------------- -> Record NEIGHBOR
  ^.* -> Error Neighbor

DESCRIPTION
  ^${SYSTEM_DESCRIPTION} -> IgnoreRemainingDescription

IgnoreRemainingDescription
  ^\S+
  ^$$ -> NEIGHBOR
  ^.* -> Error

CAPABILITY_CODES
  ^$$ -> Start
  ^\s+\(.*\)\s+\S+\,\s+\(.*\)\s+\S+\,
  ^.* -> Error CapabilityCodes`
