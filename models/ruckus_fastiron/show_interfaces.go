package ruckus_fastiron 

type ShowInterfaces struct {
	Interface	string	`json:"INTERFACE"`
	Link_status	string	`json:"LINK_STATUS"`
	Interface_up_time	string	`json:"INTERFACE_UP_TIME"`
	Protocol_status	string	`json:"PROTOCOL_STATUS"`
	Hardware_type	string	`json:"HARDWARE_TYPE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Bia	string	`json:"BIA"`
	Description	string	`json:"DESCRIPTION"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Mtu	string	`json:"MTU"`
}

var ShowInterfaces_Template = `Value Required INTERFACE (\S+)
Value LINK_STATUS (.*)
Value INTERFACE_UP_TIME (.*)
Value PROTOCOL_STATUS (.*)
Value HARDWARE_TYPE ([\w+-]+)
Value MAC_ADDRESS ([a-zA-Z0-9]+.[a-zA-Z0-9]+.[a-zA-Z0-9]+)
Value BIA ([a-zA-Z0-9]+.[a-zA-Z0-9]+.[a-zA-Z0-9]+)
Value DESCRIPTION (.*)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value PREFIX_LENGTH (\d+)
Value MTU (\d+)

Start
  ^${INTERFACE}\s+is\s+${LINK_STATUS},\s+line\s+protocol\s+is\s+${PROTOCOL_STATUS}$$
  ^\s+Port\s+(up|down)\s+for\s+${INTERFACE_UP_TIME}
  ^\s+Hardware\s+is\s+${HARDWARE_TYPE}(.*address\s+is\s+${MAC_ADDRESS})*(.*bia\s+${BIA})*
  ^\s+Description:\s+${DESCRIPTION}
  ^\s+Internet\s+address\s+is\s+${IP_ADDRESS}\/${PREFIX_LENGTH}
  ^.*MTU\s+${MTU}\.* -> Record
  ^\s+
  ^\. -> Error
`