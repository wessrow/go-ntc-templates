package models

type CiscoNxosShowInterface struct {
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