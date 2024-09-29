package cisco_ios 

type ShowIpVrfInterfaces struct {
	Interface	string	`json:"INTERFACE"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Vrf	string	`json:"VRF"`
	Proto_state	string	`json:"PROTO_STATE"`
}

var ShowIpVrfInterfaces_Template = `Value Required INTERFACE (\S+)
Value Required IP_ADDRESS (\S+)
Value Required VRF (\S+)
Value Required PROTO_STATE (\S+)

# This command returns a list of interfaces which are assigned to
# non-default VRF, including IP address (if available), VRF name,
# interface name and protocol status

Start
  ^Interface\s+IP-Address\s+VRF\s+Protocol -> Entries
  ^\s*$$
  ^. -> Error

Entries
  ^${INTERFACE}\s+${IP_ADDRESS}\s+${VRF}\s+${PROTO_STATE} -> Record
  ^\s*$$
  ^. -> Error`