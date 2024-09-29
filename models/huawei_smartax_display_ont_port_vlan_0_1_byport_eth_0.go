package models

type HuaweiSmartaxDisplayOntPortVlan01ByportEth0 struct {
	Port_id	string	`json:"PORT_ID"`
	Port_type	string	`json:"PORT_TYPE"`
	C_vlan	string	`json:"C_VLAN"`
	C_pri	string	`json:"C_PRI"`
	Eth_type	string	`json:"ETH_TYPE"`
	Vlan_type	string	`json:"VLAN_TYPE"`
	S_vlan	string	`json:"S_VLAN"`
	S_pri	string	`json:"S_PRI"`
	S_pri_policy	string	`json:"S_PRI_POLICY"`
	Native_vlan	string	`json:"NATIVE_VLAN"`
	Default_priority	string	`json:"DEFAULT_PRIORITY"`
	Downstream_mode	string	`json:"DOWNSTREAM_MODE"`
	Mismatch_policy	string	`json:"MISMATCH_POLICY"`
}

var HuaweiSmartaxDisplayOntPortVlan01ByportEth0_Template = `Value PORT_ID (\d+)
Value PORT_TYPE (\S+)
Value C_VLAN (\d+)
Value C_PRI (\d+)
Value ETH_TYPE (\S+)
Value VLAN_TYPE (\S+)
Value S_VLAN (\d+)
Value S_PRI (\d+)
Value S_PRI_POLICY (\S+)
Value Fillup NATIVE_VLAN (\d+)
Value Fillup DEFAULT_PRIORITY (\d+)
Value Fillup DOWNSTREAM_MODE (\S+)
Value Fillup MISMATCH_POLICY (\S+)

Start
  ^\s*Port\s*Port\s*C-VLAN\s*C-PRI\s*ETH-type\s*VLAN-type\s*S-VLAN\s*S-PRI\s*S-PRI\s*
  ^\s*type\s*ID\s*POLICY
  ^\s+${PORT_TYPE}\s+${PORT_ID}\s+(-|${C_VLAN})\s+(-|${C_PRI})\s+(-|${ETH_TYPE})\s+(-|${VLAN_TYPE})\s+(-|${S_VLAN})\s*(-|${S_PRI})\s*(-|${S_PRI_POLICY})\s* -> Record
  ^\s*Native\s*VLAN\s*:\s*${NATIVE_VLAN}\s*
  ^\s*Default\s*priority\s*:\s*${DEFAULT_PRIORITY}\s*
  ^\s*Downstream\s*mode\s*:\s*${DOWNSTREAM_MODE}\s*
  ^\s*Mismatch\s*policy\s*:\s*${MISMATCH_POLICY}\s* 
  ^\s*
  ^. -> Error

EOF

`