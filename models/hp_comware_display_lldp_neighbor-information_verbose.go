package models

type HpComwareDisplayLldpNeighborInformationVerbose struct {
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Chassis_id	string	`json:"CHASSIS_ID"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Neighbor_interface_description	string	`json:"NEIGHBOR_INTERFACE_DESCRIPTION"`
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Mgmt_address	string	`json:"MGMT_ADDRESS"`
	Vlan_id	string	`json:"VLAN_ID"`
}

var HpComwareDisplayLldpNeighborInformationVerbose_Template = `Value Required LOCAL_INTERFACE (\S+)
Value CHASSIS_ID (\S+)
Value NEIGHBOR_INTERFACE (.*)
Value NEIGHBOR_INTERFACE_DESCRIPTION (.*)
Value NEIGHBOR_NAME (\S+)
Value MGMT_ADDRESS (\S+)
Value VLAN_ID (\d+)

Start
  ^The\s+LLDP\s+service\s+is\s+not\s+running -> EOF
  ^LLDP\s+neighbor-information\s+of\s+port\s+\d+\[${LOCAL_INTERFACE}\]
  ^LLDP\s+agent\s+nearest-bridge:
  ^\s+LLDP\s+neighbor index\s+:.*
  ^\s+Update\s+time\s+:.*
  ^\s+Chassis\s+type\s+:.*
  ^\s+Chassis\s+ID\s+:\s+${CHASSIS_ID}
  ^\s+Port\s+ID\s+type\s+:.*
  ^\s+Port\s+ID\s+:\s+${NEIGHBOR_INTERFACE}
  ^\s+Time\s+to\s+live\d+:.*
  ^\s+Port\s+description\s+:\s+${NEIGHBOR_INTERFACE_DESCRIPTION}
  ^\s+System\s+name\s+:\s+${NEIGHBOR_NAME}
  ^\s+Management\s+address\s+:\s+${MGMT_ADDRESS}
  ^\s+Management\s+address\s+interface\s+type
  ^\s+Management\s+address\s+interface\s+ID
  ^\s+Management\s+address\s+OID
  ^\s+Port\s+VLAN\s+ID\(PVID\)\s+:\s+${VLAN_ID}
  ^\s+Maximum\s+frame\s+size -> Record Start

`