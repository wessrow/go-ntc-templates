package cisco_ios

type ShowLldpNeighborsDetail struct {
	Neighbor_name        string `json:"NEIGHBOR_NAME"`
	Capabilities         string `json:"CAPABILITIES"`
	Mgmt_address         string `json:"MGMT_ADDRESS"`
	Mac_address          string `json:"MAC_ADDRESS"`
	Serial               string `json:"SERIAL"`
	Local_interface      string `json:"LOCAL_INTERFACE"`
	Neighbor_port_id     string `json:"NEIGHBOR_PORT_ID"`
	Neighbor_interface   string `json:"NEIGHBOR_INTERFACE"`
	Manufacturer         string `json:"MANUFACTURER"`
	Sw_revision          string `json:"SW_REVISION"`
	Power_class          string `json:"POWER_CLASS"`
	Power_requested      string `json:"POWER_REQUESTED"`
	Hw_revision          string `json:"HW_REVISION"`
	Fw_revision          string `json:"FW_REVISION"`
	Power_device_type    string `json:"POWER_DEVICE_TYPE"`
	Power_priority       string `json:"POWER_PRIORITY"`
	Power_source         string `json:"POWER_SOURCE"`
	Vlan_id              string `json:"VLAN_ID"`
	Platform             string `json:"PLATFORM"`
	Power_pair           string `json:"POWER_PAIR"`
	Chassis_id           string `json:"CHASSIS_ID"`
	Neighbor_description string `json:"NEIGHBOR_DESCRIPTION"`
}

var ShowLldpNeighborsDetail_Template string = `Value LOCAL_INTERFACE (\S+)
Value CHASSIS_ID ([^:]+)
Value NEIGHBOR_PORT_ID (.*)
Value NEIGHBOR_INTERFACE (.*)
Value NEIGHBOR_NAME (.+?)
Value NEIGHBOR_DESCRIPTION (.*)
Value CAPABILITIES (.*)
Value MGMT_ADDRESS (\S+)
Value VLAN_ID (\d+)
Value MAC_ADDRESS ([0-9a-fA-F]{4}\.[0-9a-fA-F]{4}\.[0-9a-fA-F]{4})
Value SERIAL (\S+)
Value MANUFACTURER (.*)
Value HW_REVISION (\S+)
Value FW_REVISION (\S+)
Value SW_REVISION (\S+)
Value PLATFORM (.*)
Value POWER_PAIR (\S+)
Value POWER_CLASS (\d+)
Value POWER_DEVICE_TYPE (\S+)
Value POWER_PRIORITY (\S+)
Value POWER_SOURCE (\S+)
Value POWER_REQUESTED (\d+)

Start
  ^.*not advertised
  ^.*Invalid input detected -> EOF
  ^.*LLDP is not enabled -> EOF
  ^Local\s+Intf:\s+${LOCAL_INTERFACE}\s*$$
  ^Chassis\s+id:\s+${CHASSIS_ID}\s*$$
  ^Port\s+id:\s+${NEIGHBOR_PORT_ID}\s*$$
  ^Port\s+Description:\s+${NEIGHBOR_INTERFACE}\s*$$
  ^System\s+Name(\s+-\s+not\s+advertised)\s*$$
  ^System\s+Name:?\s*$$
  ^System\s+Name(:\s+${NEIGHBOR_NAME})\s*$$
  ^System\s+Description -> GetDescription
  ^Time
  ^System\s+Capabilities
  ^Enabled\s+Capabilities:\s+${CAPABILITIES}\s*$$
  ^Management\s+Addresses
  ^\s+OID
  ^\s+[\d+\.]{8,}
  ^.*IP:\s+${MGMT_ADDRESS}\s*$$
  ^Auto\s+Negotiation
  ^Physical\s+media
  # Removed \(\s+\) from the template - The line 'Other/unknown' would not be captured
  # Now looks for any text beginning with any space
  ^\s+.+\s*$$
  ^Media\s+Attachment
  ^\s+Inventory
  ^\s+Capabilities
  ^\s+Device\s+type
  ^\s+Network\s+Policies
  ^\s+Power\s+requirements
  ^\s+Location
  ^Time\s+remaining
  ^Vlan\s+ID:\s+(?:${VLAN_ID}|-\s+not\s+advertised)\s*$$
  ^Peer\s+Source\s+MAC:\s+${MAC_ADDRESS}
  ^\s+\(\S+\)
  ^(?:PoE|\s+Power)
  ^\s*-+\s*$$ -> Record
  ^Power\s+via\s+MDI -> PowerMDI
  ^MED -> Med
  ^\s*\^\s*
  ^\s*Total\s+entries\s+displayed -> Record
  ^\s*$$
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^.*$$ -> Error

GetDescription
  ^${NEIGHBOR_DESCRIPTION} -> IgnoreDescription

IgnoreDescription
  ^Time\s+remaining -> Start
  ^\S*
  ^\s*$$
  ^.* -> Error

PowerMDI
  ^\s+Power\s+Pair:\s+${POWER_PAIR}\s*$$
  ^\s+Power\s+Class:\s+${POWER_CLASS}\s*$$
  ^\s+Power\s+Device\s+Type:\s+${POWER_DEVICE_TYPE}\s*$$
  ^\s+Power\s+Source:\s+${POWER_SOURCE}\s*$$
  ^\s+Power\s+Priority:\s+${POWER_PRIORITY}\s*$$
  ^\s+Power\s+Requested:\s+${POWER_REQUESTED}\s+mW\s*$$
  ^\s*$$ -> Start

Med
  ^\s+Serial\s+number:\s+${SERIAL}\s*$$
  ^\s+Manufacturer:\s+${MANUFACTURER}\s*$$
  ^\s+Model:\s+${PLATFORM}\s*$$
  ^\s+H/W revision:\s+${HW_REVISION}\s*$$
  ^\s+F/W revision:\s+${FW_REVISION}\s*$$
  ^\s+S/W revision:\s+${SW_REVISION}\s*$$
  ^\s+\S+
  ^\s*$$
  ^\s*Total\s+entries\s+displayed -> Record
  ^\s*-+\s*$$ -> Continue.Record
  ^\s*-+\s*$$ -> Start
  ^.* -> Error
`
