package hp_procurve

type ShowLldpInfoRemoteDeviceDetail struct {
	Capabilities_supported         string `json:"CAPABILITIES_SUPPORTED"`
	Capabilities                   string `json:"CAPABILITIES"`
	Neighbor_description           string `json:"NEIGHBOR_DESCRIPTION"`
	Neighbor_interface_description string `json:"NEIGHBOR_INTERFACE_DESCRIPTION"`
	Local_interface                string `json:"LOCAL_INTERFACE"`
	Neighbor_chassis_type          string `json:"NEIGHBOR_CHASSIS_TYPE"`
	Chassis_id                     string `json:"CHASSIS_ID"`
	Port_type                      string `json:"PORT_TYPE"`
	Neighbor_interface             string `json:"NEIGHBOR_INTERFACE"`
	Neighbor_name                  string `json:"NEIGHBOR_NAME"`
	Vlan_id                        string `json:"VLAN_ID"`
	Mgmt_address                   string `json:"MGMT_ADDRESS"`
}

var ShowLldpInfoRemoteDeviceDetail_Template string = `Value Required LOCAL_INTERFACE (\S+)
Value NEIGHBOR_CHASSIS_TYPE (\S+)
Value CHASSIS_ID ([a-f0-9]{6}-[a-f0-9]{6}|(?:[a-z0-9]{2}\s){5}[a-z0-9]{2}|\S+)
Value PORT_TYPE (\S+)
Value NEIGHBOR_INTERFACE (.*?)
Value NEIGHBOR_NAME (.*?)
Value NEIGHBOR_DESCRIPTION (.*?)
Value NEIGHBOR_INTERFACE_DESCRIPTION (\S+|.*?)
Value VLAN_ID (\d+|.*?)
Value CAPABILITIES_SUPPORTED (.*?)
Value CAPABILITIES (.*?)
Value MGMT_ADDRESS (\S+|.*?)

Start
  ^\s*$$
  ^\s*LLDP\sRemote\sDevice\sInformation\sDetail -> LLDP
  ^. -> Error

LLDP
  ^\s+Local\sPort\s+:\s${LOCAL_INTERFACE}\s*$$
  ^\s+ChassisType\s+:\s${NEIGHBOR_CHASSIS_TYPE}\s*$$
  ^\s+ChassisId\s+:\s${CHASSIS_ID}\s*$$
  ^\s+PortType\s+:\s${PORT_TYPE}\s*$$
  ^\s+PortId\s+:\s${NEIGHBOR_INTERFACE}\s*$$
  ^\s+SysName\s+:\s${NEIGHBOR_NAME}\s*$$
  ^\s+System\sDescr\s:\s${NEIGHBOR_DESCRIPTION}\s*$$
  ^\s+PortDescr\s+:\s${NEIGHBOR_INTERFACE_DESCRIPTION}\s*$$
# Port VLAN ID
  ^\s+Pvid\s+:\s${VLAN_ID}\s*$$
  ^\s+System\s+Capabilities\s+Supported\s+:\s${CAPABILITIES_SUPPORTED}\s*$$
  ^\s+System\s+Capabilities\s+Enabled\s+:\s${CAPABILITIES}\s*$$
  ^\s+Remote\s+Management\s+Address\s*$$
  ^\s+Type.*$$
  ^\s+Address\s:\s${MGMT_ADDRESS}
  ^\s+Poe\s+Plus\s+Information\s+Detail\s*$$
  ^\s+Poe\s+Device\s+Type.*
  ^\s+Power\s+Source.*
  ^\s+Power\s+Priority.*
  ^\s+PD\s+Request\s+Power\s+Value.*
  ^\s+PD\s+Requested\s+Power\s+Value.*
  ^\s+Requested\s+Power\s+Value.*
  ^\s+Actual\s+Power\s+Value.*
  ^\s+PSE\s+Allocated\s+Power\s+Value.*
  ^\s+MED\sInformation\sDetail\s*$$
  ^\s+EndpointClass.*
  ^\s+Poe\s+Device\s+Type.*
  ^\s+Power\s+Requested.*
  ^\s+Power\s+Source.*
  ^\s+Power\s+Priority.*
  ^\s*$$
  ^\S+\#\s*$$ -> Record
  ^\s*-*$$ -> Record
  ^. -> Error
`
