package mikrotik_routeros 

type IpFirewallAddressListPrintTerse struct {
	Index	string	`json:"INDEX"`
	Flags	string	`json:"FLAGS"`
	List	string	`json:"LIST"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Ipv6_address	string	`json:"IPV6_ADDRESS"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Creation_time	string	`json:"CREATION_TIME"`
	Timeout	string	`json:"TIMEOUT"`
	Dynamic	string	`json:"DYNAMIC"`
}

var IpFirewallAddressListPrintTerse_Template = `Value Key INDEX (\d+)
Value FLAGS ([XD])
Value LIST (\S+)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value IPV6_ADDRESS ([a-fA-F0-9:]+)
Value PREFIX_LENGTH (\d+)
Value CREATION_TIME (\S+\s\S+)
Value TIMEOUT (\S+)
Value DYNAMIC (yes|no)

Start
  ^\s*${INDEX}\s+(${FLAGS}\s)?list=${LIST}\saddress=${IP_ADDRESS}(\/${PREFIX_LENGTH})?(\screation-time=${CREATION_TIME})?(\stimeout=${TIMEOUT})?(\sdynamic=${DYNAMIC})?.*$$ -> Record
  ^\s*${INDEX}\s+(${FLAGS}\s)?list=${LIST}\saddress=${IPV6_ADDRESS}(\/${PREFIX_LENGTH})?(\screation-time=${CREATION_TIME})?(\stimeout=${TIMEOUT})?(\sdynamic=${DYNAMIC})?.*$$ -> Record
  ^. -> Error
`