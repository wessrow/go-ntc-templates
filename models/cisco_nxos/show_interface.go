package cisco_nxos 

type ShowInterface struct {
	Interface	string	`json:"INTERFACE"`
	Link_status	string	`json:"LINK_STATUS"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Hardware_type	string	`json:"HARDWARE_TYPE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Bia	string	`json:"BIA"`
	Description	string	`json:"DESCRIPTION"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Mtu	string	`json:"MTU"`
	Mode	string	`json:"MODE"`
	Duplex	string	`json:"DUPLEX"`
	Speed	string	`json:"SPEED"`
	Input_packets	string	`json:"INPUT_PACKETS"`
	Output_packets	string	`json:"OUTPUT_PACKETS"`
	Input_errors	string	`json:"INPUT_ERRORS"`
	Output_errors	string	`json:"OUTPUT_ERRORS"`
	Bandwidth	string	`json:"BANDWIDTH"`
	Delay	string	`json:"DELAY"`
	Encapsulation	string	`json:"ENCAPSULATION"`
	Last_link_flapped	string	`json:"LAST_LINK_FLAPPED"`
	Vlan_id	string	`json:"VLAN_ID"`
	Packet_input_rate	string	`json:"PACKET_INPUT_RATE"`
	Packet_output_rate	string	`json:"PACKET_OUTPUT_RATE"`
	Bandwidth_input_rate	string	`json:"BANDWIDTH_INPUT_RATE"`
	Bandwidth_output_rate	string	`json:"BANDWIDTH_OUTPUT_RATE"`
	Media_type	string	`json:"MEDIA_TYPE"`
}

var ShowInterface_Template = `Value Required INTERFACE (\S+)
Value LINK_STATUS (.+?)
Value ADMIN_STATE (.+?)
Value HARDWARE_TYPE (.*)
Value MAC_ADDRESS ([a-zA-Z0-9]+.[a-zA-Z0-9]+.[a-zA-Z0-9]+)
Value BIA ([a-zA-Z0-9]+.[a-zA-Z0-9]+.[a-zA-Z0-9]+)
Value DESCRIPTION (\S+((\s+\S+)+)?)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value PREFIX_LENGTH (\d+)
Value MTU (\d+)
Value MODE (\S+)
Value DUPLEX (.+duplex?)
Value SPEED (.+?)
Value INPUT_PACKETS (\d+)
Value OUTPUT_PACKETS (\d+)
Value INPUT_ERRORS (\d+)
Value OUTPUT_ERRORS (\d+)
Value BANDWIDTH (\d+\s+\w+)
Value DELAY (\d+\s+\w+)
Value ENCAPSULATION ([\w\.]+)
Value LAST_LINK_FLAPPED (.+?)
Value VLAN_ID (\d+)
Value PACKET_INPUT_RATE (.+?)
Value PACKET_OUTPUT_RATE (.+?)
Value BANDWIDTH_INPUT_RATE (.+?)
Value BANDWIDTH_OUTPUT_RATE (.+?)
Value MEDIA_TYPE (.+?)

Start
  ^\S+\s+is.+ -> Continue.Record
  ^${INTERFACE}\s+is\s+${LINK_STATUS},\sline\sprotocol\sis\s${ADMIN_STATE}$$
  ^${INTERFACE}\s+is\s+${LINK_STATUS}$$
  ^admin\s+state\s+is\s+${ADMIN_STATE},
  ^\s+Hardware(:|\s+is)\s+${HARDWARE_TYPE},\s+address(:|\s+is)\s+${MAC_ADDRESS}(.*bia\s+${BIA})*
  ^\s+Description:\s+${DESCRIPTION}\s*$$
  ^\s+Internet\s+Address\s+is\s+${IP_ADDRESS}\/${PREFIX_LENGTH}
  ^\s+Port\s+mode\s+is\s+${MODE}
  ^\s+${DUPLEX}, ${SPEED}(,\s+media\s+type\s+is\s+${MEDIA_TYPE})?\s*$$
  ^\s+MTU\s+${MTU}.*BW\s+${BANDWIDTH}.*DLY\s+${DELAY}
  ^\s+Encapsulation\s+${ENCAPSULATION}(,)?(\s+Virtual\s+LAN,\s+Vlan\s+ID\s+${VLAN_ID},)?
  ^\s+${INPUT_PACKETS}\s+input\s+packets\s+\d+\s+bytes\s*$$
  ^\s+${INPUT_ERRORS}\s+input\s+error\s+\d+\s+short\s+frame\s+\d+\s+overrun\s+\d+\s+underrun\s+\d+\s+ignored\s*$$
  ^\s+${OUTPUT_PACKETS}\s+output\s+packets\s+\d+\s+bytes\s*$$
  ^\s+${OUTPUT_ERRORS}\s+output\s+error\s+\d+\s+collision\s+\d+\s+deferred\s+\d+\s+late\s+collision\s*$$
  ^\s+Last\s+link\s+flapped\s+${LAST_LINK_FLAPPED}\s*$$
  ^\s+\d+\s+seconds\s+input\s+rate\s+${BANDWIDTH_INPUT_RATE}\s+bits/sec,\s+${PACKET_INPUT_RATE}\s+packets/sec
  ^\s+\d+\s+seconds\s+output\s+rate\s+${BANDWIDTH_OUTPUT_RATE}\s+bits/sec,\s+${PACKET_OUTPUT_RATE}\s+packets/sec
`