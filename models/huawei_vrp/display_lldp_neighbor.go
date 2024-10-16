package huawei_vrp

type DisplayLldpNeighbor struct {
	Platform             string   `json:"PLATFORM"`
	Neighbor_port_id     string   `json:"NEIGHBOR_PORT_ID"`
	Neighbor_interface   string   `json:"NEIGHBOR_INTERFACE"`
	Neighbor_name        string   `json:"NEIGHBOR_NAME"`
	Vlan_id              string   `json:"VLAN_ID"`
	Serial               string   `json:"SERIAL"`
	Local_interface      string   `json:"LOCAL_INTERFACE"`
	Chassis_id           string   `json:"CHASSIS_ID"`
	Manufacturer         string   `json:"MANUFACTURER"`
	Neighbor_description []string `json:"NEIGHBOR_DESCRIPTION"`
	Capabilities         string   `json:"CAPABILITIES"`
	Mgmt_address         string   `json:"MGMT_ADDRESS"`
}

var DisplayLldpNeighbor_Template string = `Value Required LOCAL_INTERFACE (\S+)
Value CHASSIS_ID (.+?)
Value MANUFACTURER (.*?)
Value PLATFORM (.*?)
Value NEIGHBOR_PORT_ID (.*?)
Value NEIGHBOR_INTERFACE (.*?)
Value NEIGHBOR_NAME (.+?)
Value List NEIGHBOR_DESCRIPTION (.*)
Value CAPABILITIES (.*?)
Value MGMT_ADDRESS (\S+)
Value VLAN_ID (\d+)
Value SERIAL (\S+)

Start
  ^${LOCAL_INTERFACE}\s+has\s+\d+\s+neighbor(s|\(s\)):$$
  ^\S+\s+has\s+\d+\s+neighbors
  ^Neighbor\s+index
  ^Chassis\s+type
  ^Chassis\s+ID\s+:${CHASSIS_ID}$$
  ^Port\s+ID\s+type
  ^Port\s+ID\s+:${NEIGHBOR_PORT_ID}\s*$$
  ^Port\s+description\s+:${NEIGHBOR_INTERFACE}\s*$$
  ^System\s+name\s+:${NEIGHBOR_NAME}\s*$$
  ^System\s+description\s+:${NEIGHBOR_DESCRIPTION} -> SystemDescription
  ^System\s+capabilities\s+supported
  ^System\s+capabilities\s+enabled\s+:${CAPABILITIES}\s*$$
  ^Management\s+address\s+type
  ^Management\s+address\s*(value\s*)?:\s*${MGMT_ADDRESS}
  ^Expired\s+time
  ^Port\s+VLAN\s+ID\(PVID\)\s+:${VLAN_ID}
  ^VLAN\s+name\s+of\s+VLAN
  ^Protocol\s+identity
  ^Auto-negotiation
  ^OperMau
  ^Power
  ^OID
  ^PSE
  ^Port\s+power
  ^Link\s+aggregation
  ^Aggregation
  ^Maximum\s+frame\s+Size
  ^MED\s+Device\s+information -> MED
  ^\s*$$
  ^. -> Error

SystemDescription
  ^${NEIGHBOR_DESCRIPTION} -> IgnoreDescription

IgnoreDescription
  ^System\s+capabilities\s+supported -> Start
  ^${NEIGHBOR_DESCRIPTION}
  ^\s*$$
  ^.*$$ -> Error

MED
  ^Device\s+class
  ^HardwareRev
  ^FirmwareRev
  ^SoftwareRev
  ^SerialNum\s+:${SERIAL}
  ^Manufacturer\s+name\s+:${MANUFACTURER}\s*$$
  ^Model\s+name\s+:${PLATFORM}\s*$$
  ^Asset\s+tracking
  ^Media\s+policy\s+type
  ^Unknown\s+Policy
  ^VLAN\s+tagged
  ^Media
  ^Power
  ^PoE
  ^Port\s+PSE\s+Priority
  ^Port\s+Available\s+power
  ^Location\s+format
  ^Location\s+information
  ^Ca
  ^\S+\s+has\s+\d+\s+neighbors -> Record Start
  ^\s*$$
  ^. -> Error
`
