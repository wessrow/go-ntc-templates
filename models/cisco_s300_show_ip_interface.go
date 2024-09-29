package models

type CiscoS300ShowIpInterface struct {
	Ip	string	`json:"IP"`
	Interface	string	`json:"INTERFACE"`
	Interface_status_admin_oper	string	`json:"INTERFACE_STATUS_ADMIN_OPER"`
	Type	string	`json:"TYPE"`
	Status	string	`json:"STATUS"`
}

var CiscoS300ShowIpInterface_Template = `Value IP (\S+)
Value INTERFACE (\S+(?:\s\d+)?)
Value INTERFACE_STATUS_ADMIN_OPER (\S+/\S+)
Value TYPE (Static|DHCP)
Value STATUS (\S+)

Start
  ^\s*Gateway\s+IP\s+Address\s+Activity\s+status\s+Type\s*$$ -> Gateway
  ^\s*IP\s+Address\s+I/F\s+Type\s+Status\s*$$ -> IpAddress4Column
  ^\s*IP\s+Address\s+I/F\s+I/F\s+Status\s+Type\s+Status\s*$$ -> IpAddress5Column
  ^\s*$$
  ^. -> Error

Gateway
  ^(\s*-*)*\s*$$
  ^\s*IP\s+Address\s+I/F\s+Type\s+Status\s*$$ -> IpAddress4Column
  ^\s*\S+\s+\S+\s+\S+\s*$$ -> Record
  ^\s*$$
  ^. -> Error

IpAddress4Column
  ^(\s*-*)*\s*$$
  ^\s*${IP}\s+${INTERFACE}\s+${TYPE}\s+${STATUS}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

IpAddress5Column
  ^\s*admin/oper\s*$$
  ^(\s*-*)*\s*$$
  ^\s*Gateway\s+IP\s+Address\s+Activity\s+status\s+Type\s*$$ -> Gateway
  ^\s*${IP}\s+${INTERFACE}\s+${INTERFACE_STATUS_ADMIN_OPER}\s+${TYPE}\s+${STATUS}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

`