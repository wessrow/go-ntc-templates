package extreme_exos

type ShowPortsInformation struct {
	Flags       string `json:"FLAGS"`
	Link_state  string `json:"LINK_STATE"`
	Elsm        string `json:"ELSM"`
	Oam         string `json:"OAM"`
	Num_stp     string `json:"NUM_STP"`
	Num_vlan    string `json:"NUM_VLAN"`
	Interface   string `json:"INTERFACE"`
	Link_ups    string `json:"LINK_UPS"`
	Num_proto   string `json:"NUM_PROTO"`
	Jumbo_size  string `json:"JUMBO_SIZE"`
	Qos_profile string `json:"QOS_PROFILE"`
	Load_master string `json:"LOAD_MASTER"`
}

var ShowPortsInformation_Template string = `Value INTERFACE (\d+)
Value FLAGS (\S+)
Value LINK_STATE (\S+)
Value ELSM (\S+)
Value OAM (\S+)
Value LINK_UPS (\d+)
Value NUM_STP (\d+)
Value NUM_VLAN (\d+)
Value NUM_PROTO (\d+)
Value JUMBO_SIZE (\d+)
Value QOS_PROFILE (\S+)
Value LOAD_MASTER (.*)

Start
  ^\s*Port\s+Flags\s+Link\s+ELSM\s+Link\s+Num\s+Num\s+Num\s+Jumbo\s+QOS\s+Load\s*$$ -> Ports
  ^\s*
  ^. -> Error

Ports
  ^State\s+/OAM\s+UPS\s+STP\s+VLAN\s+Proto\s+Size\s+profile\s+Master\*$$
  ^(\s*=+)+\s*$$
  ^\s*${INTERFACE}\s+${FLAGS}\s+${LINK_STATE}\s+${ELSM}\s*/\s*${OAM}\s+${LINK_UPS}\s+${NUM_STP}\s+${NUM_VLAN}\s+${NUM_PROTO}\s+${JUMBO_SIZE}\s+${QOS_PROFILE}(?:\s+${LOAD_MASTER})?\s*$$ -> Record
  ^\s*>\s+indicates\s+Port\s+Display\s+Name\s+truncated\s+past\s+8\s+characters\s*$$
  ^\s*Flags\s+:\s+a\s+-\s+Load\s+Sharing\s+Algorithm\s+address-based,\s+D\s+-\s+Port\s+Disabled,\s*$$
  ^\s*e\s+-\s+Extreme\s+Discovery\s+Protocol\s+Enabled,\s+E\s+-\s+Port\s+Enabled,\s*$$
  ^\s*g\s+-\s+Egress\s+TOS\s+Enabled,\s+i\s+-\s+Isolation,\s+j\s+-\s+Jumbo\s+Frame\s+Enabled,\s*$$
  ^\s*l\s+-\s+Load\s+Sharing\s+Enabled,\s+m\s+-\s+MACLearning\s+Enabled,\s*$$
  ^\s*n\s+-\s+Ingress\s+TOS\s+Enabled,\s+o\s+-\s+Dot1p\s+Replacement\s+Enabled,\s*$$
  ^\s*P\s+-\s+Software\s+redundant\s+port\(Primary\),\s*$$
  ^\s*R\s+-\s+Software\s+redundant\s+port\(Redundant\),\s*$$
  ^\s*q\s+-\s+Background\s+QOS\s+Monitoring\s+Enabled,\s*$$
  ^\s*s\s+-\s+diffserv\s+Replacement\s+Enabled,\s*$$
  ^\s*v\s+-\s+Vman\s+Enabled,\s+f\s+-\s+Unicast\s+Flooding\s+Enabled,\s*$$
  ^\s*M\s+-\s+Multicast\s+Flooding\s+Enabled,\s+B\s+-\s+Broadcast\s+Flooding\s+Enabled\s*$$
  ^\s*L\s+-\s+Extreme\s+Link\s+Status\s+Monitoring\s+Enabled\s*$$
  ^\s*O\s+-\s+Ethernet\s+OAM\s+Enabled\s*$$
  ^\s*w\s+-\s+MACLearning\s+Disabled\s+with\s+Forwarding\s*$$
  ^\s*b\s+-\s+Rx\s+and\s+Tx\s+Flow\s+Control\s+Enabled,\s+x\s+-\s+Rx\s+Flow\s+Control\s+Enabled\s*$$
  ^\s*p\s+-\s+Priority\s+Flow\s+Control\s+Enabled\s*$$
  ^\s*
  ^. -> Error
`
