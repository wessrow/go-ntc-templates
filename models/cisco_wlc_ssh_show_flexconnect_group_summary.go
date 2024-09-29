package models

type CiscoWlcSshShowFlexconnectGroupSummary struct {
	Flexconnect_group_count	string	`json:"FLEXCONNECT_GROUP_COUNT"`
	Flexconnect_group_name	string	`json:"FLEXCONNECT_GROUP_NAME"`
	Ap_count	string	`json:"AP_COUNT"`
}

var CiscoWlcSshShowFlexconnectGroupSummary_Template = `Value Filldown FLEXCONNECT_GROUP_COUNT (\d+)
Value Required FLEXCONNECT_GROUP_NAME (.*\S)
Value AP_COUNT (\d+)


Start
  ^\s*FlexConnect\sGroup\s+Summary:\s*Count:\s+${FLEXCONNECT_GROUP_COUNT}
  ^\s*Group\s+Name\s+#\s+Aps -> Flexconnect
  ^\s*$$
  ^. -> Error

Flexconnect
  ^-+
  ^${FLEXCONNECT_GROUP_NAME}\s+${AP_COUNT} -> Record
  ^\s*$$
  ^. -> Error

`