package cisco_wlc 

type SshShowInterfaceSummary struct {
	Int_count	string	`json:"INT_COUNT"`
	Name	string	`json:"NAME"`
	Port	string	`json:"PORT"`
	Vlan_id	string	`json:"VLAN_ID"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Type	string	`json:"TYPE"`
	Ap_mgr	string	`json:"AP_MGR"`
	Guest	string	`json:"GUEST"`
}

var SshShowInterfaceSummary_Template = `Value Filldown INT_COUNT (\d+)
Value Required NAME (\S+)
Value PORT (\S+)
Value VLAN_ID (\S+)
Value IP_ADDRESS (([\d1-9]+\.?){4})
Value TYPE (\S+)
Value AP_MGR (\S+)
Value GUEST (\S+)

Start
  ^\s*Number\sof\sInterfaces\.*\s${INT_COUNT}s*$$
  ^Interface\s+Name\s+Port\s+Vlan\s+Id\s+IP\s+Address\s+Type\s+Ap\s+Mgr\s+Guest -> Type_One
  ^\s*$$
  ^. -> Error

Type_One
  ^-+\s
  ^${NAME}\s\s+${PORT}\s+${VLAN_ID}\s+${IP_ADDRESS}\s+${TYPE}\s+${AP_MGR}\s+${GUEST} -> Record
  ^\s*$$
  ^. -> Error
`