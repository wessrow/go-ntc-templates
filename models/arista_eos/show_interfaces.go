package arista_eos

type ShowInterfaces struct {
	Link_status_change string `json:"LINK_STATUS_CHANGE"`
	Interface          string `json:"INTERFACE"`
	Link_status        string `json:"LINK_STATUS"`
	Bia                string `json:"BIA"`
	Ip_address         string `json:"IP_ADDRESS"`
	Mtu                string `json:"MTU"`
	Interface_up_time  string `json:"INTERFACE_UP_TIME"`
	Protocol_status    string `json:"PROTOCOL_STATUS"`
	Hardware_type      string `json:"HARDWARE_TYPE"`
	Mac_address        string `json:"MAC_ADDRESS"`
	Description        string `json:"DESCRIPTION"`
	Bandwidth          string `json:"BANDWIDTH"`
}

var ShowInterfaces_Template string = `Value Required INTERFACE (\S+)
Value LINK_STATUS (.*)
Value PROTOCOL_STATUS (.*)
Value HARDWARE_TYPE ([\w+-]+)
Value MAC_ADDRESS ([a-zA-Z0-9]+.[a-zA-Z0-9]+.[a-zA-Z0-9]+)
Value BIA ([a-zA-Z0-9]+.[a-zA-Z0-9]+.[a-zA-Z0-9]+)
Value DESCRIPTION ([^\"]*)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+\/\d+)
Value MTU (\d+)
Value BANDWIDTH (\d+\s+\w+)
Value INTERFACE_UP_TIME (.*)
Value LINK_STATUS_CHANGE (\d+)

Start
  ^\S+\s+is\s+\S+,\s+line\s+protocol\s+is\s+\S+\s+\S+ -> Continue.Record
  ^${INTERFACE}\s+is\s+${LINK_STATUS},\s+line\s+protocol\s+is\s+${PROTOCOL_STATUS}
  ^\s+Hardware\s+is\s+${HARDWARE_TYPE}(.*address\s+is\s+${MAC_ADDRESS})*(.*bia\s+${BIA})*
  ^\s+Description:\s+"?${DESCRIPTION}"?$$
  ^\s+Internet\s+address\s+is\s+${IP_ADDRESS}
  ^\s+(Up|Down)\s+${INTERFACE_UP_TIME}
  ^\s+${LINK_STATUS_CHANGE}\s+link\s+status\s+changes.*
  ^.*MTU\s+${MTU}(.*BW\s+${BANDWIDTH})*
  ^\s*
  ^. -> Error

EOF
`
