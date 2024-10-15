package huawei_smartax

type DisplayOntPortVlan01Byvlan0 struct {
	Port_id      string `json:"PORT_ID"`
	S_pri        string `json:"S_PRI"`
	Port_type    string `json:"PORT_TYPE"`
	C_pri        string `json:"C_PRI"`
	Eth_type     string `json:"ETH_TYPE"`
	Vlan_type    string `json:"VLAN_TYPE"`
	S_vlan       string `json:"S_VLAN"`
	S_pri_policy string `json:"S_PRI_POLICY"`
	C_vlan       string `json:"C_VLAN"`
}

var DisplayOntPortVlan01Byvlan0_Template string = `Value C_VLAN (\d+)
Value C_PRI (\d+)
Value ETH_TYPE (\S+)
Value VLAN_TYPE (\S+)
Value PORT_TYPE (\S+)
Value PORT_ID (\d+)
Value S_VLAN (\d+)
Value S_PRI (\d+)
Value S_PRI_POLICY (\S+)

Start
  ^\s*C-VLAN\s*C-PRI\s*ETH-type\s*VLAN-type\s*Port\s*Port\s*S-VLAN\s*S-PRI\s*
  ^\s*type\s*ID\s*POLICY
  ^\s+${C_VLAN}\s+(-|${C_PRI})\s+${ETH_TYPE}\s+${VLAN_TYPE}\s+${PORT_TYPE}\s+${PORT_ID}\s+${S_VLAN}\s+(-|${S_PRI})\s*(-|${S_PRI_POLICY})\s* -> Record
  ^\s*
  ^. -> Error
`
