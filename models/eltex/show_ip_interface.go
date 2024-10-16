package eltex

type ShowIpInterface struct {
	Ip                     string `json:"IP"`
	Interface_status_oper  string `json:"INTERFACE_STATUS_OPER"`
	Type                   string `json:"TYPE"`
	Directed_broadcast     string `json:"DIRECTED_BROADCAST"`
	Precedence             string `json:"PRECEDENCE"`
	Redirect               string `json:"REDIRECT"`
	Status                 string `json:"STATUS"`
	Gateway_type           string `json:"GATEWAY_TYPE"`
	Gateway_status         string `json:"GATEWAY_STATUS"`
	Interface              string `json:"INTERFACE"`
	Interface_status_admin string `json:"INTERFACE_STATUS_ADMIN"`
	Gateway_ip             string `json:"GATEWAY_IP"`
}

var ShowIpInterface_Template string = `Value GATEWAY_IP (\S+)
Value GATEWAY_STATUS (\S+)
Value GATEWAY_TYPE (\S+)
Value IP (\S+)
Value INTERFACE (.*?)
Value INTERFACE_STATUS_ADMIN (\S+)
Value INTERFACE_STATUS_OPER (\S+)
Value TYPE (Static|DHCP)
Value DIRECTED_BROADCAST (disable|enable)
Value PRECEDENCE (Yes|No)
Value REDIRECT (enable|disable)
Value STATUS (\S+)

Start
  ^\s*Gateway\s+IP\s+Address\s+Activity\s+status\s+Type\s*$$ -> Gateway
  ^\s*IP\s+Address\s+I/F\s+Type\s+Status\s*$$ -> Column4
  ^\s*IP\s+Address\s+I/F\s+Type\s+Directed\s+Precedence\s+Status\s*$$ -> Column6
  ^\s*IP\s+Address\s+I/F\s+I/F\s+Status\s+Type\s+Directed\s+Prec\s+Redirect\s+Status\s*$$ -> Column8
  ^\s*$$
  ^. -> Error

Gateway
  ^(?:\s*-+)+\s*$$
  ^${GATEWAY_IP}\s+${GATEWAY_STATUS}\s+${GATEWAY_TYPE}\s*$$ -> Record
  ^\s*IP\s+Address\s+I/F\s+Type\s+Status\s*$$ -> Column4
  ^\s*$$
  ^. -> Error

Column4
  ^(?:\s*-+)+\s*$$
  ^${IP}\s+${INTERFACE}\s+${TYPE}\s+${STATUS}\s*$$ -> Record
  # Skip wrapped lines
  ^\s+received\s*$$
  ^\s+receiv\s*$$
  ^\s+ed\s*$$
  ^\s*$$
  ^. -> Error

Column6
  ^(?:\s*-+)+\s*$$
  ^\s+Broadcast\s*$$
  ^${IP}\s+${INTERFACE}\s+${TYPE}\s+${DIRECTED_BROADCAST}\s+${PRECEDENCE}\s+${STATUS}\s*$$ -> Record
  # Skip wrapped lines
  ^\s+received\s*$$
  ^\s+receiv\s*$$
  ^\s+ed\s*$$
  ^\s*$$
  ^. -> Error

Column8
  ^\s+admin/oper\s+Broadcast\s*$$
  ^(?:\s*-+)+\s*$$
  ^${IP}\s+${INTERFACE}(?:\s+${INTERFACE_STATUS_ADMIN}/${INTERFACE_STATUS_OPER})?\s+${TYPE}\s+${DIRECTED_BROADCAST}\s+${PRECEDENCE}\s+${REDIRECT}\s+${STATUS}\s*$$ -> Record
  # Skip wrapped lines
  ^\s+received\s*$$
  ^\s+receiv\s*$$
  ^\s+ed\s*$$
  ^\s*$$
  ^. -> Error
`
