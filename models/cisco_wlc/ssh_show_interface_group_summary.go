package cisco_wlc 

type SshShowInterfaceGroupSummary struct {
	Interface_group_name	string	`json:"INTERFACE_GROUP_NAME"`
	Total_interfaces	string	`json:"TOTAL_INTERFACES"`
	Total_wlans	string	`json:"TOTAL_WLANS"`
	Total_ap_groups	string	`json:"TOTAL_AP_GROUPS"`
}

var SshShowInterfaceGroupSummary_Template = `Value INTERFACE_GROUP_NAME (\S+)
Value TOTAL_INTERFACES (\d+)
Value TOTAL_WLANS (\d+)
Value TOTAL_AP_GROUPS (\d+)

Start
  ^Interface\s*Group\s*Name\s*Total\s*Interfaces\s*Total\s*Wlans\s*Total\s*AP\s*Groups\s*Quarantine
  ^-+
  ^${INTERFACE_GROUP_NAME}\s+${TOTAL_INTERFACES}\s+${TOTAL_WLANS}\s+${TOTAL_AP_GROUPS} -> Record
  ^\s*$$
  ^. -> Error
`