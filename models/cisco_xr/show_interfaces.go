package cisco_xr 

type ShowInterfaces struct {
	Interface	string	`json:"INTERFACE"`
	Link_status	string	`json:"LINK_STATUS"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Hardware_type	string	`json:"HARDWARE_TYPE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Bia	string	`json:"BIA"`
	Description	string	`json:"DESCRIPTION"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Mtu	string	`json:"MTU"`
	Duplex	string	`json:"DUPLEX"`
	Hardware_media	string	`json:"HARDWARE_MEDIA"`
	Speed	string	`json:"SPEED"`
	Bandwidth	string	`json:"BANDWIDTH"`
	Encapsulation	string	`json:"ENCAPSULATION"`
	Vlan_id	string	`json:"VLAN_ID"`
	Last_input	string	`json:"LAST_INPUT"`
	Last_output	string	`json:"LAST_OUTPUT"`
	Input_rate	string	`json:"INPUT_RATE"`
	Output_rate	string	`json:"OUTPUT_RATE"`
	Input_pps	string	`json:"INPUT_PPS"`
	Output_pps	string	`json:"OUTPUT_PPS"`
	Input_packets	string	`json:"INPUT_PACKETS"`
	Output_packets	string	`json:"OUTPUT_PACKETS"`
	Runts	string	`json:"RUNTS"`
	Giants	string	`json:"GIANTS"`
	Input_errors	string	`json:"INPUT_ERRORS"`
	Crc	string	`json:"CRC"`
	Frame	string	`json:"FRAME"`
	Overrun	string	`json:"OVERRUN"`
	Abort	string	`json:"ABORT"`
	Output_errors	string	`json:"OUTPUT_ERRORS"`
}

var ShowInterfaces_Template = `Value Required INTERFACE (\S+)
Value LINK_STATUS (.+?)
Value ADMIN_STATE (.+?)
Value HARDWARE_TYPE (\S+?(?:\s+Ethernet|))
Value MAC_ADDRESS ((?:\w{4}\.){2}\w{4})
Value BIA ((?:\w{4}\.){2}\w{4})
Value DESCRIPTION (.*?)
Value IP_ADDRESS (.*?)
Value MTU (\d+)
Value DUPLEX (.+?)
Value HARDWARE_MEDIA (.*?)
Value SPEED (.+?b/s)
Value BANDWIDTH (\d+\s+\w+)
Value ENCAPSULATION ([\w\.\s]+)
Value VLAN_ID (\d+)
Value LAST_INPUT (.+?)
Value LAST_OUTPUT (.+?)
Value INPUT_RATE (\d+)
Value OUTPUT_RATE (\d+)
Value INPUT_PPS (\d+)
Value OUTPUT_PPS (\d+)
Value INPUT_PACKETS (\d+)
Value OUTPUT_PACKETS (\d+)
Value RUNTS (\d+)
Value GIANTS (\d+)
Value INPUT_ERRORS (\d+)
Value CRC (\d+)
Value FRAME (\d+)
Value OVERRUN (\d+)
Value ABORT (\d+)
Value OUTPUT_ERRORS (\d+)

Start
  ^\S+\s+is -> Continue.Record
  ^${INTERFACE}\sis\s+${LINK_STATUS},\s+line\sprotocol\sis\s+${ADMIN_STATE}\s*$$
  ^\s+Hardware\s+is\s+${HARDWARE_TYPE}(?:\s+(sub-)?interface\(s\)|)(?:,\s+address\s+is\s+${MAC_ADDRESS}(?:\s+\(bia\s+${BIA}\)\s*)*$$|\s.+|\s*$$)
  ^\s+Description:\s+${DESCRIPTION}\s*$$
  ^\s+[Ii]nternet\s+[Aa]ddress\s+is\s+${IP_ADDRESS}\s*$$
  ^\s+MTU\s+${MTU}.*BW\s+${BANDWIDTH}
  ^\s+Encapsulation\s+${ENCAPSULATION},\s+(Vlan|VLAN)\s+I(D|d)\s+${VLAN_ID}.+$$
  ^\s+Encapsulation\s+${ENCAPSULATION},.*$$
  ^\s+(?:[Dd]uplex\s+|)${DUPLEX}(?:-[Dd]uplex|),\s+${SPEED}(?:,\s+${HARDWARE_MEDIA},)?
  ^\s+Last\s+input\s+${LAST_INPUT},\s+output\s+${LAST_OUTPUT}\s*$$
  ^.*input\s+rate\s+${INPUT_RATE}\s+\w+/sec,\s+${INPUT_PPS}\s+packets.+$$
  ^.*output\s+rate\s+${OUTPUT_RATE}\s+\w+/sec,\s+${OUTPUT_PPS}\s+packets.+$$
  ^\s+${INPUT_PACKETS}\s+packets\s+input,\s+\d+\s+bytes,\s+\d+(\s+\S+){3}\s*$$
  ^\s+\d+\s+drops\s+for\s+unrecognized(\s+\S+)+$$
  ^\s+Received\s+\d+
  ^\s+${RUNTS}\s+runts,\s+${GIANTS}\s+giants,\s+\d+\s+throttles,\s+\d+\s+parity\s*$$
  ^\s+${INPUT_ERRORS}\s+input\s+errors,\s+${CRC}\s+CRC,\s+${FRAME}\s+frame,\s+${OVERRUN}\s+overrun,\s+\d+\s+ignored\s*$$
  ^\s+${INPUT_ERRORS}\s+input\s+errors,\s+${CRC}\s+CRC,\s+${FRAME}\s+frame,\s+${OVERRUN}\s+overrun,\s+\d+\s+ignored,\s+${ABORT}\s+abort\s*$$
  ^\s+${OUTPUT_PACKETS}\s+packets\s+output,\s+\d+\s+bytes,\s+\d+(\s+\S+){3}\s*$$
  ^\s+Output\s+\d+
  ^\s+${OUTPUT_ERRORS}\s+output\s+errors,(\s+\d+\s+\S+,?)+\s*$$
  ^\s+\d+output\s+buffer
  ^\s+\d+carrier\s+transitions
`