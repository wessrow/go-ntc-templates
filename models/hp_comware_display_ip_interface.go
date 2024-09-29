package models

type HpComwareDisplayIpInterface struct {
	Interface	string	`json:"INTERFACE"`
	Line_status	string	`json:"LINE_STATUS"`
	Protocol_status	string	`json:"PROTOCOL_STATUS"`
	Route_map	string	`json:"ROUTE_MAP"`
	Ip_address	[]string	`json:"IP_ADDRESS"`
	Mtu	string	`json:"MTU"`
}

var HpComwareDisplayIpInterface_Template = `Value INTERFACE (\S+)
Value LINE_STATUS (UP|DOWN|Administratively DOWN)
Value PROTOCOL_STATUS (UP(\(spoofing\))?|DOWN)
Value ROUTE_MAP (\S+)
Value List IP_ADDRESS (\S+)
Value MTU (\d+)


Start
  ^${INTERFACE}\s+current\s+state\s*:\s*${LINE_STATUS} -> Interface
  ^. -> Error


Interface
  ^\S+\s+current\s+state -> Continue.Record
  ^${INTERFACE}\s+current\s+state\s*:\s*${LINE_STATUS}
  ^Line\s+protocol\s+current\s+state\s*:\s*${PROTOCOL_STATUS}
  ^Internet\s+Address\s+is\s+${IP_ADDRESS}\s+Primary
  ^Internet\s+Address\s+is\s+${IP_ADDRESS}\s+Sub
  ^Broadcast\s+address\s*:\s*\S+
  ^The\s+Maximum\s+Transmit\s+Unit\s*:\s*${MTU}\s+bytes
  ^Policy\s+routing\s+is\s+enabled,\s+using\s+route\s+map\s+${ROUTE_MAP}
  ^input\spackets\s*:\s*\d+,\sbytes\s*:\s*\d+,\smulticasts\s*:\s*\d+
  ^output\spackets\s*:\s*\d+,\sbytes\s*:\s*\d+,\smulticasts\s*:\s*\d+
  ^TTL\sinvalid\spacket\snumber\s*:\s*\d+
  ^ICMP\spacket\sinput\snumber\s*:\s*\d+
  ^ARP\spacket\sinput\snumber\s*:\s*\d+
  ^\s+Echo\sreply\s*:\s*\d+
  ^\s+Unreachable\s*:\s*\d+
  ^\s+Source\squench\s*:\s*\d+
  ^\s+Routing\sredirect\s*:\s*\d+
  ^\s+Echo\srequest\s*:\s*\d+
  ^\s+Router\sadvert\s*:\s*\d+
  ^\s+Router\ssolicit\s*:\s*\d+
  ^\s+Time\sexceed\s*:\s*\d+
  ^\s+IP\sheader\sbad\s*:\s*\d+
  ^\s+Timestamp\srequest\s*:\s*\d+
  ^\s+Timestamp\sreply\s*:\s*\d+
  ^\s+Information\srequest\s*:\s*\d+
  ^\s+Information\sreply\s*:\s*\d+
  ^\s+Netmask\srequest\s*:\s*\d+
  ^\s+Netmask\sreply\s*:\s*\d+
  ^\s+Unknown\stype\s*:\s*\d+
  ^\s+Request\spacket\s*:\s*\d+
  ^\s+Reply\spacket\s*:\s*\d+
  ^\s+Unknown\spacket\s*:\s*\d+
  ^\s*$$
  ^. -> Error

`