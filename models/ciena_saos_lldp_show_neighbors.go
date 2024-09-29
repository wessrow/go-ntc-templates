package models

type CienaSaosLldpShowNeighbors struct {
	Remote_ipv4_address	string	`json:"REMOTE_IPV4_ADDRESS"`
	Remote_ipv6_address	string	`json:"REMOTE_IPV6_ADDRESS"`
	System_name	string	`json:"SYSTEM_NAME"`
	System_description	string	`json:"SYSTEM_DESCRIPTION"`
	Local_port	string	`json:"LOCAL_PORT"`
	Remote_port	string	`json:"REMOTE_PORT"`
	Remote_chassis_id	string	`json:"REMOTE_CHASSIS_ID"`
	Remote_mgmt_ipv4_address	[]string	`json:"REMOTE_MGMT_IPV4_ADDRESS"`
	Remote_mgmt_ipv6_address	[]string	`json:"REMOTE_MGMT_IPV6_ADDRESS"`
	Remote_system_name	string	`json:"REMOTE_SYSTEM_NAME"`
	Remote_system_description	string	`json:"REMOTE_SYSTEM_DESCRIPTION"`
	Remote_synce_support	string	`json:"REMOTE_SYNCE_SUPPORT"`
	Remote_synce_config	string	`json:"REMOTE_SYNCE_CONFIG"`
}

var CienaSaosLldpShowNeighbors_Template = `Value List,Filldown REMOTE_IPV4_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value List,Filldown REMOTE_IPV6_ADDRESS ([a-f0-9:]+)
Value Filldown SYSTEM_NAME (\S+)
Value Filldown SYSTEM_DESCRIPTION (.*\S)
Value Required LOCAL_PORT (\S+)
Value REMOTE_PORT (\S+)
Value REMOTE_CHASSIS_ID (\S+)
Value List REMOTE_MGMT_IPV4_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value List REMOTE_MGMT_IPV6_ADDRESS ([a-f0-9:]+)
Value REMOTE_SYSTEM_NAME (\S+)
Value REMOTE_SYSTEM_DESCRIPTION (.*\S)
Value REMOTE_SYNCE_SUPPORT (.*\S)
Value REMOTE_SYNCE_CONFIG (.*\S)

Start
  ^\+\-+
  ^\|\s*LLDP\s+Neighbors\s+\|\s*$$
  ^\|\s*Local\s+\|\s+Remote\s+\|\s*$$
  ^\|\s*Port\s+\|\s*Port\s+\|\s*Info\s+\|\s*$$
  ^\|\sRemote\sAddr:\s${REMOTE_IPV4_ADDRESS}
  ^\|\sRemote\sAddr:\s${REMOTE_IPV6_ADDRESS}
  ^\|\sSystem\sName:\s${SYSTEM_NAME}\s+\|
  ^\|\sSystem\sDesc:\s${SYSTEM_DESCRIPTION}\s+\|
  # neighbor record
  ^\|${LOCAL_PORT}\s*\|${REMOTE_PORT}\s*\|\s*Chassis\sId:\s*${REMOTE_CHASSIS_ID} -> Neighbor
  ^\s*$$
  ^. -> Error

Neighbor
  ^\|\s*\|\s*\|\s*Mgmt\sAddr:\s${REMOTE_MGMT_IPV4_ADDRESS}\s+\| -> NeighborMgmt
  ^\|\s*\|\s*\|\s*Mgmt\sAddr:\s${REMOTE_MGMT_IPV6_ADDRESS}\s+\| -> NeighborMgmt
  ^\|\s*\|\s*\|\s*System\sName:\s${REMOTE_SYSTEM_NAME} -> Neighbor
  ^\|\s*\|\s*\|\s*System\sDesc:\s${REMOTE_SYSTEM_DESCRIPTION}\s+\|
  ^\|\s*\|\s*\|\s*SyncE\sSuppt:\s${REMOTE_SYNCE_SUPPORT}\s+\|
  ^\|\s*\|\s*\|\s*SyncE\sConfg:\s${REMOTE_SYNCE_CONFIG}\s+\| 
  ^\+\-+ -> Record Start
  ^\s*$$
  ^. -> Error

NeighborMgmt
  ^\|\s*\|\s*\|\s+${REMOTE_MGMT_IPV4_ADDRESS}\s+\|$$
  ^\|\s*\|\s*\|\s+${REMOTE_MGMT_IPV6_ADDRESS}\s+\|$$
  ^\|\s*\|\s*\|\s*System\sName:\s${REMOTE_SYSTEM_NAME} -> Neighbor
  ^\s*$$
  ^. -> Error

`