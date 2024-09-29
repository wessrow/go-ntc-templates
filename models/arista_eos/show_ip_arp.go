package arista_eos 

type ShowIpArp struct {
	Ip_address	string	`json:"IP_ADDRESS"`
	Age	string	`json:"AGE"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Interface	string	`json:"INTERFACE"`
	Vrf	string	`json:"VRF"`
}

var ShowIpArp_Template = `Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value AGE (\S+)
Value MAC_ADDRESS (\S+)
Value INTERFACE (.*)
Value Filldown VRF (\S+)

Start
  ^VRF:\s+${VRF}
  ^Address
  ^${IP_ADDRESS}\s+${AGE}\s+${MAC_ADDRESS}\s+${INTERFACE} -> Record
  ^. -> Error

EOF
`