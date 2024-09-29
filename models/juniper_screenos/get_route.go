package juniper_screenos 

type GetRoute struct {
	Vr	string	`json:"VR"`
	Best	string	`json:"BEST"`
	Id	string	`json:"ID"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Outinterface	string	`json:"OUTINTERFACE"`
	Nexthop	string	`json:"NEXTHOP"`
	Protocol	string	`json:"PROTOCOL"`
	Pref	string	`json:"PREF"`
	Metric	string	`json:"METRIC"`
	Vsys	string	`json:"VSYS"`
}

var GetRoute_Template = `Value Filldown VR (\S+)
Value BEST (\S+)
Value ID (\d+)
Value Required IP_ADDRESS (\d+(\.\d+){3})
Value Required PREFIX_LENGTH (\d{1,2})
Value OUTINTERFACE (\S+)
Value NEXTHOP (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value PROTOCOL (\S+)
Value PREF (\d+)
Value METRIC (\d+)
Value VSYS (\S+)

Start
  # Match the VR
  ^IPv\d\s+Dest-Routes\s+for\s+<${VR}>
  # Match route line
  ^${BEST}\s+${ID}\s+${IP_ADDRESS}\/${PREFIX_LENGTH}\s+${OUTINTERFACE}\s+${NEXTHOP}\s+${PROTOCOL}\s+${PREF}\s+${METRIC}\s+${VSYS} -> Record
  # Match table separator
  ^-+\s*$$
  # Match route type keys
  ^\S+:\s+
  # Match lines that are just space
  ^\s+$$
  # Match table header
  ^\s+ID\s+IP-Prefix\s+Interface\s+Gateway\s+P\s+Pref\s+Mtr\s+Vsys\s*$$
  # Error out if line does not match any expressions
  ^.+ -> Error "Line not found"

EOF
`