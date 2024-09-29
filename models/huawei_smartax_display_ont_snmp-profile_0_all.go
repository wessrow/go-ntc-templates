package models

type HuaweiSmartaxDisplayOntSnmpProfile0All struct {
	Ont_id	string	`json:"ONT_ID"`
	Snmp_profile_id	string	`json:"SNMP_PROFILE_ID"`
	Snmp_profile_name	string	`json:"SNMP_PROFILE_NAME"`
}

var HuaweiSmartaxDisplayOntSnmpProfile0All_Template = `Value ONT_ID (\d+)
Value SNMP_PROFILE_ID (\d+)
Value SNMP_PROFILE_NAME (\S+)

Start
  ^\s*ONT\s*ID\s*SNMP\s*profile\s*ID\s*SNMP\s*profile\s*name\s*
  ^\s*${ONT_ID}\s*${SNMP_PROFILE_ID}\s*${SNMP_PROFILE_NAME}
  ^\s*
  ^. -> Error

`