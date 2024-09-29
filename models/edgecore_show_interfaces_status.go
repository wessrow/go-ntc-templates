package models

type EdgecoreShowInterfacesStatus struct {
	Port	string	`json:"PORT"`
	Port_type	string	`json:"PORT_TYPE"`
	Address	string	`json:"ADDRESS"`
	Name	string	`json:"NAME"`
	Port_admin	string	`json:"PORT_ADMIN"`
	Mdix_mode	string	`json:"MDIX_MODE"`
	Negotiation	string	`json:"NEGOTIATION"`
	Speed	string	`json:"SPEED"`
	Duplex	string	`json:"DUPLEX"`
	Capabilities	string	`json:"CAPABILITIES"`
	Broadcast_storm	string	`json:"BROADCAST_STORM"`
	Broadcast_storm_limit	string	`json:"BROADCAST_STORM_LIMIT"`
	Multicast_storm	string	`json:"MULTICAST_STORM"`
	Multicast_storm_limit	string	`json:"MULTICAST_STORM_LIMIT"`
	Unknownunicast_storm	string	`json:"UNKNOWNUNICAST_STORM"`
	Unknownunicast_storm_limit	string	`json:"UNKNOWNUNICAST_STORM_LIMIT"`
	Flow_control	string	`json:"FLOW_CONTROL"`
	Vlan_trunking	string	`json:"VLAN_TRUNKING"`
	Lacp	string	`json:"LACP"`
	Port_security	string	`json:"PORT_SECURITY"`
	Max_mac_count	string	`json:"MAX_MAC_COUNT"`
	Port_security_action	string	`json:"PORT_SECURITY_ACTION"`
	Media_type	string	`json:"MEDIA_TYPE"`
	Giga_phy_mode	string	`json:"GIGA_PHY_MODE"`
	Link_status	string	`json:"LINK_STATUS"`
	Link_down_reason	string	`json:"LINK_DOWN_REASON"`
	Operation_status	string	`json:"OPERATION_STATUS"`
	Port_uptime	string	`json:"PORT_UPTIME"`
	Flow_control_type	string	`json:"FLOW_CONTROL_TYPE"`
}

var EdgecoreShowInterfacesStatus_Template = `Value Required PORT (.*)
Value PORT_TYPE (.*)
Value ADDRESS (.*)
Value NAME (.*)
Value PORT_ADMIN (.*)
Value MDIX_MODE (.*)
Value NEGOTIATION (.*?)
Value SPEED (\d+)
Value DUPLEX (\S+)
Value CAPABILITIES (.*)
Value BROADCAST_STORM (.*)
Value BROADCAST_STORM_LIMIT (.*)
Value MULTICAST_STORM (.*)
Value MULTICAST_STORM_LIMIT (.*)
Value UNKNOWNUNICAST_STORM (.*)
Value UNKNOWNUNICAST_STORM_LIMIT (.*)
Value FLOW_CONTROL (.*)
Value VLAN_TRUNKING (.*)
Value LACP (.*)
Value PORT_SECURITY (.*)
Value MAX_MAC_COUNT (.*)
Value PORT_SECURITY_ACTION (.*)
Value MEDIA_TYPE (.*)
Value GIGA_PHY_MODE (.*)
Value LINK_STATUS (.*)
Value LINK_DOWN_REASON (.*)
Value OPERATION_STATUS (.*)
Value PORT_UPTIME (.*)
Value FLOW_CONTROL_TYPE (.*)

Start
  ^\s*Information\s+of.*$$ -> Continue.Record
  ^\s*Information\s+of\s+${PORT}\s*$$
  ^\s*M(ac|AC)\s+Address:\s*${ADDRESS}\s*$$
  ^\s*Basic\s+Information:\s*$$ -> BasicInfo
  ^\s*$$
  ^. -> Error

BasicInfo
  ^\s*Port\s+Type:\s*${PORT_TYPE}\s*$$
  ^\s*M(ac|AC)\s+Address:\s*${ADDRESS}\s*$$
  ^\s*Configuration:\s*$$ -> Configuration
  ^\s*$$
  ^. -> Error

Configuration
  ^\s*Name:\s*${NAME}\s*$$
  ^\s*Port\s+Admin:\s*${PORT_ADMIN}\s*$$
  ^\s*MDIX\s+mode:\s*${MDIX_MODE}\s*$$
  ^\s*Speed-duplex:\s*${NEGOTIATION}\s*$$
  ^\s*Capabilities:\s*${CAPABILITIES}\s*$$
  ^\s*Broadcast\s+Storm:\s*${BROADCAST_STORM}\s*$$
  ^\s*Broadcast\s+Storm\s+Limit:\s*${BROADCAST_STORM_LIMIT}\s*$$
  ^\s*Multicast\s+Storm:\s*${MULTICAST_STORM}\s*$$
  ^\s*Multicast\s+Storm\s+Limit:\s*${MULTICAST_STORM_LIMIT}\s*$$
  ^\s*UnknownUnicast\s+Storm:\s*${UNKNOWNUNICAST_STORM}\s*$$
  ^\s*UnknownUnicast\s+Storm\s+Limit:\s*${UNKNOWNUNICAST_STORM_LIMIT}\s*$$
  ^\s*Flow\s+Control:\s*${FLOW_CONTROL}\s*$$
  ^\s*VLAN\s+Trunking:\s*${VLAN_TRUNKING}\s*$$
  ^\s*LACP:\s*${LACP}\s*$$
  ^\s*Port\s+Security:\s*${PORT_SECURITY}\s*$$
  ^\s*Max\s+MAC\s+Count:\s*${MAX_MAC_COUNT}\s*$$
  ^\s*Port\s+Security\s+Action:\s*${PORT_SECURITY_ACTION}\s*$$
  ^\s*Media\s+Type:\s*${MEDIA_TYPE}\s*$$
  ^\s*Giga\s+PHY\s+mode:\s*${GIGA_PHY_MODE}\s*$$
  ^\s*Current Status:\s*$$ -> CurrentStatus
  ^\s*$$
  ^. -> Error

CurrentStatus
  ^\s*Link\s+Status:\s*${LINK_STATUS}\s*$$
  ^\s*Link\s+Down\s+Reason:\s*${LINK_DOWN_REASON}\s*$$
  ^\s*Port\s+Operation\s+Status:\s*${OPERATION_STATUS}\s*$$
  ^\s*Operation\s+Speed-duplex:\s*${SPEED}${DUPLEX}\s*$$
  ^\s*Port\s+Uptime:\s*${PORT_UPTIME}\s*$$
  ^\s*Flow\s+Control\s+Type:\s*${FLOW_CONTROL_TYPE}\s*$$
  ^\s*$$ -> Start
  ^. -> Error

`