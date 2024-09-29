package zte_zxros 

type ShowInterface struct {
	Interface	string	`json:"INTERFACE"`
	Link_status	string	`json:"LINK_STATUS"`
	Interface_index	string	`json:"INTERFACE_INDEX"`
	Description	string	`json:"DESCRIPTION"`
	Protocol_status	string	`json:"PROTOCOL_STATUS"`
	Ipv4_protocol	string	`json:"IPV4_PROTOCOL"`
	Ipv6_protocol	string	`json:"IPV6_PROTOCOL"`
	Detect_status	string	`json:"DETECT_STATUS"`
	Uptime	string	`json:"UPTIME"`
	Hardware_type	string	`json:"HARDWARE_TYPE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Mask	string	`json:"MASK"`
	Bandwith	string	`json:"BANDWITH"`
	Mtu	[]string	`json:"MTU"`
}

var ShowInterface_Template = `Value Required INTERFACE (\S+)
Value LINK_STATUS (.+?)
Value INTERFACE_INDEX (\d+)
Value DESCRIPTION (.*?)
Value PROTOCOL_STATUS (.+?)
Value IPV4_PROTOCOL (.+?)
Value IPV6_PROTOCOL (.+?)
Value DETECT_STATUS (\S+)
Value UPTIME (.+?)
Value HARDWARE_TYPE (.+?)
Value MAC_ADDRESS (\S+)
Value IP_ADDRESS (\S+?)
Value MASK (\d*)
Value BANDWITH (.+?)
Value List MTU (\d+)

Start
  ^\S -> Continue.Record
  ^${INTERFACE}\s+is\s+${LINK_STATUS},\s+ifindex:\s+${INTERFACE_INDEX}\s*$$
  ^\s+Description:\s+${DESCRIPTION}\s*$$
  ^\s+Line\s+protocol\s+is\s+${PROTOCOL_STATUS},\s+IPv4\s+protocol\s+is\s+${IPV4_PROTOCOL},\s+IPv6\s+protocol\s+is\s+${IPV6_PROTOCOL},\s*$$
  ^\s+detected\s+status\s+is\s+${DETECT_STATUS}\s*$$
  ^\s+Last\s+line\s+protocol\s+up\s+time\s+:\s+${UPTIME}\s*$$
  ^\s+Hardware\s+is\s+${HARDWARE_TYPE},\s+address\s+is\s+${MAC_ADDRESS}\s*$$
  ^\s+Internet\s+address\s+is\s+${IP_ADDRESS}/?${MASK}\s*$$
  ^\s+BW\s+${BANDWITH}\s*$$
  ^\s+IP\s+MTU\s+${MTU}\s+bytes
  ^\s+MTU\s+${MTU}\s+bytes
  ^\s+IPv6\s+MTU\s+${MTU}\s+bytes
  ^\s+MPLS\s+MTU\s+${MTU}\s+bytes
  ^\s+The\s+interface\s+is
  ^\s+Fec-(eth|bypass)\s+:
  ^\s+ARP\s+(type|Timeout)
  ^\s+Last\s+Clear\s+Time
  ^\s+Rate\s+period
  ^\s+Input\s+:
  ^\s+Output\s+:
  ^\s+Peak\s+rate:
  ^\s+Intf\s+utilization:
  ^\s+HardwareCounters:                                          
  ^\s+In_Bytes
  ^\s+In_Broadcasts
  ^\s+In_Unicasts
  ^\s+In_Fragments
  ^\s+In_65_127B
  ^\s+In_256_511B
  ^\s+In_1024_1518B
  ^\s+In_Undersize
  ^\s+E_Bytes
  ^\s+E_Broadcasts
  ^\s+E_Unicasts
  ^\s+E_64B
  ^\s+E_128_255B
  ^\s+E_512_1023B
  ^\s+E_1519_MaxB
  ^\s+StreamCounters
  ^\s*$$
  ^. -> Error
`