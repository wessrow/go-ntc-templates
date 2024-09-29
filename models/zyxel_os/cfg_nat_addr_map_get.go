package zyxel_os 

type CfgNatAddrMapGet struct {
	Index	string	`json:"INDEX"`
	Name	string	`json:"NAME"`
	Interface	string	`json:"INTERFACE"`
	Type	string	`json:"TYPE"`
	Local_start_ip	string	`json:"LOCAL_START_IP"`
	Local_end_ip	string	`json:"LOCAL_END_IP"`
	Global_start_ip	string	`json:"GLOBAL_START_IP"`
	Global_end_ip	string	`json:"GLOBAL_END_IP"`
}

var CfgNatAddrMapGet_Template = `Value INDEX (\d+)
Value NAME (.+\w)
Value INTERFACE ((Default|WWAN|(A|V)?DSL))
Value TYPE (([oO]ne|[mM]any)-to-([oO]ne|[mM]any))
Value LOCAL_START_IP (\d+\.\d+\.\d+\.\d+)
Value LOCAL_END_IP ((\d+\.\d+\.\d+\.\d+)?)
Value GLOBAL_START_IP (\d+\.\d+\.\d+\.\d+)
Value GLOBAL_END_IP ((\d+\.\d+\.\d+\.\d+)?)

Start
  ^Index\s+MappingRuleName\s+Interface\s+Type\s+LocalStartIP\s+LocalEndIP\s+GlobalStartIP\s+GlobalEndIP\s*$$ -> NAT_ADDRTable
  ^\s*$$
  ^. -> Error

NAT_ADDRTable
  ^${INDEX}\s+${NAME}\s+${INTERFACE}\s+${TYPE}\s+${LOCAL_START_IP}\s+${LOCAL_END_IP}\s+${GLOBAL_START_IP}\s+${GLOBAL_END_IP}\s*$$ -> Record
  ^Command\sSuccessful.\s*$$
  ^\s*$$
  ^. -> Error
`